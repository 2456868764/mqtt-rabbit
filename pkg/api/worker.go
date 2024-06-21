package api

import (
	"strconv"
	"time"

	"bifromq_engine/pkg/db"
	"bifromq_engine/pkg/model/entity"
	"bifromq_engine/pkg/model/req"
	"bifromq_engine/pkg/model/resp"
	"github.com/gin-gonic/gin"
)

var Worker = &worker{}

type worker struct {
}

func (worker) LoadConfigurations(c *gin.Context) {
	var params req.ConfigurationReq
	err := c.Bind(&params)
	if err != nil {
		Error(c, 20001, err.Error())
		return
	}
	data := &resp.ConfigurationResp{
		Data: make(map[string]string),
	}
	lastUpdateTime := time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)
	var configs []entity.DataSource
	db.DB.Model(entity.DataSource{}).Where("confType = ? and status = 1 order by id ASC", params.ConfType).Find(&configs)
	for _, config := range configs {
		data.Data[config.Name] = config.Content
		if config.UpdateTime.After(lastUpdateTime) {
			lastUpdateTime = config.UpdateTime
		}
	}
	data.LastUpdateTime = lastUpdateTime
	Success(c, data)
}

func (worker) Delete(c *gin.Context) {
	// uid := c.Param("id")
	//err := db.DB.Transaction(func(tx *gorm.DB) error {
	//	tx.Where("id =?", uid).Delete(&entity.User{})
	//	tx.Where("userId =?", uid).Delete(&entity.UserRolesRole{})
	//	tx.Where("userId =?", uid).Delete(&entity.Profile{})
	//	return nil
	//})
	//if err != nil {
	//	Error(c, 20001, err.Error())
	//	return
	//}
	Success(c, "")
}

func (worker) Register(c *gin.Context) {
	var params req.WorkerRegisterReq
	var err error
	err = c.Bind(&params)
	if err != nil {
		Error(c, 20001, err.Error())
		return
	}
	var newWorker = entity.Worker{
		Name:            params.Name,
		IP:              params.IP,
		Tag:             params.Tag,
		Port:            params.Port,
		Status:          int32(entity.WorkerStatusRegister),
		HeartbeatMisses: 0,
		CreateTime:      time.Now(),
		UpdateTime:      time.Now(),
	}
	orm := db.DB.Model(entity.Worker{})
	exsitedWorker := entity.Worker{}

	err = orm.Where("name =?", params.Name).Find(&exsitedWorker).Error
	if err != nil {
		Error(c, 500, err.Error())
		return
	}

	if exsitedWorker.ID > 0 {
		//existed
		newWorker.ID = exsitedWorker.ID
		err = orm.Where("id = ?", exsitedWorker.ID).Updates(newWorker).Error
	} else {
		err = orm.Create(&newWorker).Error
	}
	if err != nil {
		Error(c, 500, err.Error())
		return
	}

	data := &resp.WorkerRegisterResp{
		WorkerID: newWorker.ID,
	}
	Success(c, data)
}

func (worker) UnRegister(c *gin.Context) {
	var params req.WorkerUnRegisterReq
	var err error
	err = c.Bind(&params)
	if err != nil {
		Error(c, 20001, err.Error())
		return
	}
	orm := db.DB.Model(entity.Worker{})
	err = orm.Where("id = ?", params.WorkerID).Update("status", int32(entity.WorkerStatusUnRegister)).Error
	if err != nil {
		Error(c, 500, err.Error())
		return
	}
	Success(c, "")
}

func (worker) Heartbeat(c *gin.Context) {
	var params req.WorkerHeartbeatReq
	var err error
	err = c.Bind(&params)
	if err != nil {
		Error(c, 20001, err.Error())
		return
	}

	existedWorker := entity.Worker{}
	orm := db.DB.Model(entity.Worker{})
	err = orm.Where("id = ?", params.WorkerID).Find(&existedWorker).Error
	if err != nil {
		Error(c, 500, err.Error())
		return
	}
	if existedWorker.ID == 0 {
		Error(c, 20001, "worker not found")
		return
	}

	// Get last datasource time
	lastSourcesTime := time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)
	lastSinksTime := time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)
	var sourceCount int64
	err = db.DB.Model(entity.DataSource{}).Where("confType = ? AND status = 1", entity.ConfTypeSource).Count(&sourceCount).Error
	if err != nil {
		Error(c, 500, err.Error())
		return
	}
	if sourceCount == 0 {
		lastSourcesTime = time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC) // Unix 纪元时间
	} else {
		err = db.DB.Model(entity.DataSource{}).Where("confType = ? AND status = 1", entity.ConfTypeSource).Select("MAX(updateTime)").Scan(&lastSourcesTime).Error
		if err != nil {
			Error(c, 500, err.Error())
			return
		}
	}

	var sinkCount int64
	err = db.DB.Model(entity.DataSource{}).Where("confType = ? AND status = 1", entity.ConfTypeSink).Count(&sinkCount).Error
	if err != nil {
		Error(c, 500, err.Error())
		return
	}
	if sinkCount == 0 {
		lastSinksTime = time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC) // Unix 纪元时间
	} else {
		err = db.DB.Model(entity.DataSource{}).Where("confType = ? AND status = 1", entity.ConfTypeSink).Select("MAX(updateTime)").Scan(&lastSinksTime).Error
		if err != nil {
			Error(c, 500, err.Error())
			return
		}
	}

	// Check worker status
	if existedWorker.Status == int32(entity.WorkerStatusUnRegister) || existedWorker.Status == int32(entity.WorkerStatusHeartbeatMiss) || existedWorker.Status == int32(entity.WorkerStatusDeleted) {
		data := &resp.WorkerHeartbeatResp{
			LastSourcesTime: lastSourcesTime,
			LastSinksTime:   lastSinksTime,
			Kill:            true,
		}
		Success(c, data)
		return
	}

	// Update worker heartbeat
	upWorker := entity.Worker{
		IP:              params.IP,
		Port:            params.Port,
		HeartbeatMisses: 0,
		HeartbeatTime:   time.Now(),
		Status:          int32(entity.WorkerStatusOK),
		LastSourcesTime: params.LastSourcesTime,
	}
	err = orm.Where("id = ?", params.WorkerID).Updates(upWorker).Error
	if err != nil {
		Error(c, 500, err.Error())
		return
	}

	data := &resp.WorkerHeartbeatResp{
		LastSourcesTime: lastSourcesTime,
		LastSinksTime:   lastSinksTime,
		Kill:            false,
	}
	Success(c, data)
}

func (worker) List(c *gin.Context) {
	var data = resp.WorkerListResp{
		PageData: make([]resp.WorkerListItem, 0),
	}
	var status = c.DefaultQuery("status", "")
	var pageNoReq = c.DefaultQuery("pageNo", "1")
	var pageSizeReq = c.DefaultQuery("pageSize", "10")
	pageNo, _ := strconv.Atoi(pageNoReq)
	pageSize, _ := strconv.Atoi(pageSizeReq)
	var workerList []entity.Worker
	orm := db.DB.Model(entity.Worker{})
	if status != "" {
		if statusInt, err := strconv.Atoi(status); err == nil {
			if statusInt > 0 {
				orm = orm.Where("status=?", statusInt)
			}
		}
	}
	orm.Count(&data.Total)
	orm.Offset((pageNo - 1) * pageSize).Limit(pageSize).Find(&workerList)
	for _, datum := range workerList {
		workerListItem := resp.WorkerListItem{}
		workerListItem.ID = datum.ID
		workerListItem.Name = datum.Name
		workerListItem.IP = datum.IP
		workerListItem.Port = datum.Port
		workerListItem.LastSourcesTime = datum.LastSourcesTime
		workerListItem.Status = datum.Status
		workerListItem.StatusText = entity.WorkerStatus_name[datum.Status]
		workerListItem.HeartbeatTime = datum.HeartbeatTime
		workerListItem.HeartbeatMisses = datum.HeartbeatMisses
		workerListItem.Tag = datum.Tag
		workerListItem.CreateTime = datum.CreateTime
		workerListItem.UpdateTime = datum.UpdateTime
		data.PageData = append(data.PageData, workerListItem)
	}
	Success(c, data)
}

func (worker) GetConfiguration(c *gin.Context) {
	var name = c.DefaultQuery("name", "")
	if len(name) == 0 {
		Error(c, 20001, "name is empty")
		return
	}
	data := &resp.GetConfigurationResp{}
	var config entity.DataSource
	db.DB.Model(entity.DataSource{}).Where("name = ? and status = 1", name).Find(&config)
	data.Content = config.Content
	data.Name = name
	Success(c, data)
}

func (worker) UpdateConfiguration(c *gin.Context) {
	var params req.ConfigurationUpdateReq
	var err error
	err = c.Bind(&params)
	if err != nil {
		Error(c, 20001, err.Error())
		return
	}
	updateConfig := &entity.DataSource{
		Content:    params.Content,
		UpdateTime: time.Now(),
	}
	db.DB.Model(entity.DataSource{}).Where("name = ?", params.Name).Updates(updateConfig)
	Success(c, "")
}
