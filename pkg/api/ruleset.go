package api

import (
	"bifromq_engine/pkg/errcode"
	"crypto/md5"
	"fmt"
	"strconv"
	"time"

	"bifromq_engine/pkg/db"
	"bifromq_engine/pkg/model/entity"
	"bifromq_engine/pkg/model/req"
	"bifromq_engine/pkg/model/resp"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var RuleSet = &ruleSet{}

type ruleSet struct {
}

func (ruleSet) Detail(c *gin.Context) {
	var data = &resp.RuleSetResp{
		Streams: make([]*resp.StreamListItem, 0),
		Rules:   make([]*resp.RuleListItem, 0),
	}

	var ruleSetID, _ = c.Get("id")
	db.DB.Model(entity.RuleSet{}).Where("id = ? and deleted = 0", ruleSetID).Find(&data)

	if data.WorkerID > 0 {
		var node entity.Worker
		db.DB.Model(entity.Worker{}).Where("id = ?", data.WorkerID).First(&node)
		data.Node = &resp.RuleSetNode{
			ID:     node.ID,
			Name:   node.Name,
			IP:     node.IP,
			Port:   node.Port,
			Status: node.Status,
			Tag:    node.Tag,
		}
	}
	var streams []entity.Stream
	db.DB.Model(entity.Stream{}).Where("ruleSetID = ? and deleted = 0", ruleSetID).Find(&streams)
	for _, stream := range streams {
		streamListItem := resp.StreamListItem{
			Stream: stream,
		}
		data.Streams = append(data.Streams, &streamListItem)
	}
	var rules []entity.Rule
	db.DB.Model(entity.Rule{}).Where("ruleSetID = ? and deleted = 0", ruleSetID).Find(&rules)
	for _, rule := range rules {
		ruleListItem := resp.RuleListItem{
			Rule:            rule,
			StatusText:      entity.TaskStatus_name[rule.Status],
			StatusCheckDesc: entity.TaskCheckStatus_name[rule.StatusCheck],
		}
		data.Rules = append(data.Rules, &ruleListItem)
	}
	Success(c, data)

}

func (ruleSet) List(c *gin.Context) {
	var data = resp.RuleSetListResp{
		PageData: make([]resp.RuleSetListItem, 0),
	}
	var status = c.DefaultQuery("status", "")
	var pageNoReq = c.DefaultQuery("pageNo", "1")
	var pageSizeReq = c.DefaultQuery("pageSize", "10")
	pageNo, _ := strconv.Atoi(pageNoReq)
	pageSize, _ := strconv.Atoi(pageSizeReq)
	var ruleSetList []entity.RuleSet
	orm := db.DB.Model(entity.RuleSet{})
	if status != "" {
		if statusInt, err := strconv.Atoi(status); err == nil {
			if statusInt > 0 {
				orm = orm.Where("status = ? and deleted  = ?", statusInt, 0)
			}
		}
	} else {
		orm = orm.Where("deleted  = ? ", 0)
	}

	orm.Count(&data.Total)
	orm.Offset((pageNo - 1) * pageSize).Limit(pageSize).Find(&ruleSetList)
	for _, ruleSet := range ruleSetList {
		ruleSetItem := resp.RuleSetListItem{
			StreamCount: 0,
			RuleCount:   0,
		}
		// get stream count
		var err error
		var streamCount int64
		err = db.DB.Model(entity.Stream{}).Where("ruleSetID = ? and deleted = 0 ", ruleSet.ID).Count(&streamCount).Error
		if err != nil {
			Error(c, 500, err.Error())
			return
		}
		ruleSetItem.StreamCount = int(streamCount)
		// rule count
		var ruleCount int64
		err = db.DB.Model(entity.Rule{}).Where("ruleSetID = ?  and deleted = 0 ", ruleSet.ID).Count(&ruleCount).Error
		if err != nil {
			Error(c, 500, err.Error())
			return
		}
		ruleSetItem.RuleCount = int(ruleCount)
		ruleSetItem.StatusText = entity.TaskStatus_name[ruleSet.Status]
		ruleSetItem.StatusCheckText = entity.TaskCheckStatus_name[ruleSet.StatusCheck]
		// get node info
		if ruleSet.WorkerID > 0 {
			var node entity.Worker
			db.DB.Model(entity.Worker{}).Where("id=?", ruleSet.WorkerID).First(&node)
			ruleSetItem.Node = &resp.RuleSetNode{
				ID:     node.ID,
				Name:   node.Name,
				IP:     node.IP,
				Port:   node.Port,
				Status: node.Status,
				Tag:    node.Tag,
			}
		}
		ruleSetItem.ID = ruleSet.ID
		ruleSetItem.Name = ruleSet.Name
		ruleSetItem.Status = ruleSet.Status
		ruleSetItem.StatusCheck = ruleSet.StatusCheck
		ruleSetItem.StatusCheckTime = ruleSet.StatusCheckTime
		ruleSetItem.ScheduleTime = ruleSet.ScheduleTime
		ruleSetItem.UpdateTime = ruleSet.UpdateTime
		ruleSetItem.CreateTime = ruleSet.CreateTime
		data.PageData = append(data.PageData, ruleSetItem)
	}
	Success(c, data)
}

func (ruleSet) Update(c *gin.Context) {
	var params req.PatchUserReq
	err := c.BindJSON(&params)
	if err != nil {
		Error(c, 20001, err.Error())
		return
	}
	orm := db.DB.Model(entity.User{}).Where("id=?", params.Id)
	if params.Password != nil {
		orm.Update("password", fmt.Sprintf("%x", md5.Sum([]byte(*params.Password))))
	}
	if params.Enable != nil {
		orm.Update("enable", *params.Enable)
	}
	if params.Username != nil {
		orm.Update("username", *params.Username)
		db.DB.Model(entity.Profile{}).Where("userId=?", params.Id).Update("nickName", *params.Username)
	}
	if params.RoleIds != nil {
		db.DB.Where("userId=?", params.Id).Delete(entity.UserRolesRole{})
		if len(*params.RoleIds) > 0 {
			for _, i2 := range *params.RoleIds {
				db.DB.Model(entity.UserRolesRole{}).Create(&entity.UserRolesRole{
					UserId: params.Id,
					RoleId: i2,
				})
			}
		}
	}

	Success(c, err)
}

func (ruleSet) Add(c *gin.Context) {
	var params req.RuleSetAddReq
	var err error
	err = c.Bind(&params)
	if err != nil {
		Error(c, 20001, err.Error())
		return
	}
	var newRuleSet = entity.RuleSet{
		Name:       params.Name,
		Status:     int32(entity.TaskStatus_INIT),
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
		WorkerID:   0,
	}

	orm := db.DB.Model(entity.RuleSet{})
	exsitedRuleSet := entity.RuleSet{}
	err = orm.Where("name =?", params.Name).Find(&exsitedRuleSet).Error
	if err != nil {
		Error(c, 500, err.Error())
		return
	}

	if exsitedRuleSet.ID > 0 {
		ErrorStatus(c, errcode.RuleExistedError)
		return
	}
	orm.Create(&newRuleSet)
	Success(c, "")
}

func (ruleSet) Delete(c *gin.Context) {
	uid := c.Param("id")
	err := db.DB.Transaction(func(tx *gorm.DB) error {
		tx.Where("id =?", uid).Delete(&entity.User{})
		tx.Where("userId =?", uid).Delete(&entity.UserRolesRole{})
		tx.Where("userId =?", uid).Delete(&entity.Profile{})
		return nil
	})
	if err != nil {
		Error(c, 20001, err.Error())
		return
	}
	Success(c, "")
}
