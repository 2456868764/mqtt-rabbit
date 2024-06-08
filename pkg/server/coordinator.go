package server

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"bifromq_engine/pkg/agent"
	"bifromq_engine/pkg/api"
	"bifromq_engine/pkg/db"
	"bifromq_engine/pkg/logs"
	"bifromq_engine/pkg/model/entity"
	"bifromq_engine/pkg/model/req"
	"bifromq_engine/pkg/routes"
	"github.com/gin-gonic/gin"
)

const (
	defaultHeartbeat         = 10 * time.Second
	shutdownTimeout          = 5 * time.Second
	defaultMaxMisses         = 3
	defaultScheduleInterval  = 10 * time.Second
	defaultTaskCheckInterval = 30 * time.Second
)

type Coordinator struct {
	gin                  *gin.Engine
	maxHeartbeatMisses   uint8
	heartbeatInterval    time.Duration
	taskScheduleInterval time.Duration
	taskCheckInterval    time.Duration
	coordinatorPort      int
	stopChan             <-chan struct{}
}

func InitCoordinator(coordinatorPort int) (*Coordinator, error) {
	r := gin.Default()
	coordinator := &Coordinator{
		gin:                  r,
		coordinatorPort:      coordinatorPort,
		maxHeartbeatMisses:   defaultMaxMisses,
		heartbeatInterval:    defaultHeartbeat,
		taskScheduleInterval: defaultScheduleInterval,
		taskCheckInterval:    defaultTaskCheckInterval,
	}
	// init routes
	// Register middlewares
	r.Use(routes.Logger())
	// Register routes
	r.POST("/api/configuration/load", api.Worker.LoadConfigurations)
	r.POST("/api/node/register", api.Worker.Register)
	r.POST("/api/node/unregister", api.Worker.UnRegister)
	r.POST("/api/node/heartbeat", api.Worker.Heartbeat)
	return coordinator, nil
}

func (c *Coordinator) Run(stopChan <-chan struct{}) error {
	//
	go c.gin.Run(fmt.Sprintf(":%d", c.coordinatorPort))
	//go c.startWorkerCheck(stopChan)
	go c.startTaskSchedule(stopChan)
	go c.startTaskCheck(stopChan)
	return nil
}

func (c *Coordinator) startTaskSchedule(stopChan <-chan struct{}) error {
	ticker := time.NewTicker(c.taskScheduleInterval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			c.scheduleTasks()
		case <-stopChan:
			return nil
		}
	}
}

func (c *Coordinator) startTaskCheck(stopChan <-chan struct{}) error {
	ticker := time.NewTicker(c.taskCheckInterval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			c.checkTasks()
		case <-stopChan:
			return nil
		}
	}
}

func (c *Coordinator) startWorkerCheck(stopChan <-chan struct{}) error {
	ticker := time.NewTicker(c.heartbeatInterval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			c.removeInactiveWorkers()
		case <-stopChan:
			return nil
		}
	}
}

func (c *Coordinator) removeInactiveWorkers() {
	logs.Infof("[scheduler] removeInactiveWorkers running")
	var workerList []entity.Worker
	db.DB.Model(entity.Worker{}).Where("status = ?", entity.WorkerStatusOK).Order("`id` Asc").Find(&workerList)
	for _, worker := range workerList {
		if worker.HeartbeatMisses > int32(c.maxHeartbeatMisses) {
			db.DB.Model(entity.Worker{}).Where("id = ?", worker.ID).Update("status", entity.WorkerStatusHeartbeatMiss)
		} else {
			upWorker := entity.Worker{
				HeartbeatMisses: worker.HeartbeatMisses + 1,
			}
			db.DB.Model(entity.Worker{}).Where("id = ?", worker.ID).Updates(upWorker)
		}
	}
}

func (c *Coordinator) scheduleTasks() {
	logs.Infof("[scheduler] scheduleTasks running")
	var workerList []entity.Worker
	db.DB.Model(entity.Worker{}).Where("status = ?", entity.WorkerStatusOK).Order("`id` Asc").Find(&workerList)
	if len(workerList) == 0 {
		return
	}
	logs.Infof("[scheduler] get worker count=%d", len(workerList))
	// get ruleset
	var ruleSets []entity.RuleSet
	db.DB.Model(entity.RuleSet{}).Where("status = ?", entity.TaskStatus_QUEUED).Order("`id` Asc").Find(&ruleSets)
	logs.Infof("[scheduler] get ruleset count=%d", len(ruleSets))
	for _, ruleSet := range ruleSets {
		// get streams and rules
		var streams []entity.Stream
		var rules []entity.Rule
		db.DB.Model(entity.Stream{}).Where("ruleSetID = ?", ruleSet.ID).Order("`id` Asc").Find(&streams)
		db.DB.Model(entity.Rule{}).Where("ruleSetID = ?", ruleSet.ID).Order("`id` Asc").Find(&rules)

		ruleSetsReq := req.EkuiperRuleSetReq{
			Streams: make(map[string]string),
			Rules:   make(map[string]string),
			Tables:  make(map[string]string),
		}
		for _, stream := range streams {
			ruleSetsReq.Streams[stream.Name] = stream.Statement
		}
		for _, rule := range rules {
			item := req.EkuiperRuleSetRuleItemReq{}
			item.ID = rule.Name
			if rule.RuleType == string(entity.RuleTypeGraph) {
				item.Graph = rule.Statement
			} else {
				item.Sql = rule.Statement
			}
			actions := make([]map[string]interface{}, 0)
			json.Unmarshal([]byte(rule.Actions), &actions)
			item.Actions = actions
			bytes, _ := json.Marshal(item)
			ruleSetsReq.Rules[rule.Name] = string(bytes)
		}

		logs.Infof("[schedule] ruleSetsReq = %+v", ruleSetsReq)
		bytes, _ := json.Marshal(ruleSetsReq)
		importReq := req.EkuiperRuleSetImportReq{
			Content: string(bytes),
		}
		shuffle(workerList)
		index := 0
		for {
			worker := workerList[index]
			url := fmt.Sprintf("http://%s:%d", worker.IP, worker.Port)
			logs.Infof("[schedule] worker id=%d url=%s", worker.ID, url)
			client := agent.NewClient(url)
			bytes, _ = json.Marshal(importReq)
			reqBody := string(bytes)
			if err := client.RulesetImport(reqBody); err == nil {
				// update
				updateRuleSet := entity.RuleSet{
					Status:       int(entity.TaskStatus_STARTED),
					WorkerID:     worker.ID,
					ScheduleTime: time.Now(),
				}
				db.DB.Model(entity.RuleSet{}).Where("id = ?", ruleSet.ID).Updates(updateRuleSet)
				db.DB.Model(entity.Stream{}).Where("ruleSetID = ?", ruleSet.ID).Update("status", int(entity.TaskStatus_STARTED))
				db.DB.Model(entity.Rule{}).Where("ruleSetID = ?", ruleSet.ID).Update("status", int(entity.TaskStatus_STARTED))
				break
			} else {
				logs.Errorf("[schedule] schedule ruleset id %d, call worker server id:%d, url:%s import ruleSet error:%v", ruleSet.ID, worker.ID, url, err)
			}
			index++
			if index == len(workerList) {
				break
			}
		}

	}
}

func (c *Coordinator) checkTasks() {
	// get ruleset
	logs.Infof("[scheduler] checkTasks running")
	var ruleSets []entity.RuleSet
	db.DB.Model(entity.RuleSet{}).Where("status = ?", entity.TaskStatus_STARTED).Order("`id` Asc").Find(&ruleSets)
	for _, ruleSet := range ruleSets {
		// get streams and rules
		//var streams []entity.Stream
		var worker entity.Worker
		db.DB.Model(entity.Worker{}).Where("id = ?", ruleSet.WorkerID).Find(&worker)
		if worker.ID == 0 {
			logs.Errorf("can not find worker for ruleset %v", ruleSet)
			continue
		}
		url := fmt.Sprintf("http://%s:%d", worker.IP, worker.Port)
		client := agent.NewClient(url)
		var rules []entity.Rule
		//db.DB.Model(entity.Stream{}).Where("ruleSetID = ?", ruleSet.ID).Order("`id` Asc").Find(&streams)
		db.DB.Model(entity.Rule{}).Where("ruleSetID = ?", ruleSet.ID).Order("`id` Asc").Find(&rules)
		totalSuccess := 0
		for _, rule := range rules {
			status, err := client.RuleStatus(rule.Name)
			if err != nil {
				// update rule status
				updateRule := entity.Rule{
					StatusCheck:     int(entity.TaskCheckStatus_FAILED),
					StatusCheckText: err.Error(),
					StatusCheckTime: time.Now(),
				}
				db.DB.Model(entity.Rule{}).Where("id = ?", rule.ID).Updates(updateRule)
			} else {
				result := make(map[string]string)
				json.Unmarshal([]byte(status), &result)
				statusCheck := entity.TaskCheckStatus_STOPPED
				if result["status"] == "running" {
					statusCheck = entity.TaskCheckStatus_RUNNING
					totalSuccess++
				}
				// update rule status
				updateRule := entity.Rule{
					StatusCheck:     int(statusCheck),
					StatusCheckText: status,
					StatusCheckTime: time.Now(),
				}
				db.DB.Model(entity.Rule{}).Where("id = ?", rule.ID).Updates(updateRule)
			}
		}
		totalCheckStatus := entity.TaskCheckStatus_RUNNING
		if totalSuccess == 0 {
			totalCheckStatus = entity.TaskCheckStatus_FAILED
		} else {
			if totalSuccess < len(rules) {
				totalCheckStatus = entity.TaskCheckStatus_PARTFAILED
			}
		}
		// update ruleSet status
		updateRuleSet := entity.RuleSet{
			StatusCheck:     int(totalCheckStatus),
			StatusCheckTime: time.Now(),
		}
		db.DB.Model(entity.RuleSet{}).Where("id = ?", ruleSet.ID).Updates(updateRuleSet)
	}
}

func shuffle(arr []entity.Worker) {
	rand.Seed(time.Now().UnixNano()) // 设置随机种子
	for i := len(arr) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)           // 生成0到i之间的随机数
		arr[i], arr[j] = arr[j], arr[i] // 交换位置
	}
}
