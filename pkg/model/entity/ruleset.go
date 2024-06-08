package entity

import "time"

type TaskStatus int

const (
	TaskStatus_INIT     TaskStatus = 0
	TaskStatus_QUEUED   TaskStatus = 1
	TaskStatus_STARTED  TaskStatus = 2
	TaskStatus_COMPLETE TaskStatus = 3
	TaskStatus_STOPPED  TaskStatus = 4
	TaskStatus_FAILED   TaskStatus = 5
)

type TaskCheckStatus int

const (
	TaskCheckStatus_INIT       TaskCheckStatus = 0
	TaskCheckStatus_RUNNING    TaskCheckStatus = 1
	TaskCheckStatus_STOPPED    TaskCheckStatus = 2
	TaskCheckStatus_FAILED     TaskCheckStatus = 3
	TaskCheckStatus_PARTFAILED TaskCheckStatus = 4
)

// Enum value maps for TaskStatus.
var (
	TaskStatus_name = map[int32]string{
		0: "QUEUED",
		1: "STARTED",
		2: "COMPLETE",
		3: "FAILED",
	}
	TaskStatus_value = map[string]int32{
		"QUEUED":   0,
		"STARTED":  1,
		"COMPLETE": 2,
		"FAILED":   3,
	}
)

type RuleType string

const (
	RuleTypeSQL   RuleType = "sql"
	RuleTypeGraph RuleType = "graph"
)

type RuleSet struct {
	ID              int32     `json:"id"`
	Name            string    `json:"name"`
	Status          int       `json:"status"`
	WorkerID        int32     `json:"WorkerID"  gorm:"column:workerID"`
	StatusCheck     int       `json:"statusCheck" gorm:"column:statusCheck"`
	StatusCheckTime time.Time `json:"statusCheckTime" gorm:"column:statusCheckTime"`
	ScheduleTime    time.Time `json:"scheduleTime" gorm:"column:scheduleTime"`
	CreateTime      time.Time `json:"createTime" gorm:"column:createTime"`
	UpdateTime      time.Time `json:"updateTime" gorm:"column:updateTime"`
}

func (RuleSet) TableName() string {
	return "rule_set"
}

type Stream struct {
	ID         int32     `json:"id"`
	RuleSetID  int32     `json:"ruleSetID"  gorm:"column:ruleSetID"`
	Name       string    `json:"name"`
	Statement  string    `json:"statement"`
	Status     int       `json:"status"`
	CreateTime time.Time `json:"createTime" gorm:"column:createTime"`
	UpdateTime time.Time `json:"updateTime" gorm:"column:updateTime"`
}

func (Stream) TableName() string {
	return "stream"
}

type Rule struct {
	ID              int32     `json:"id"`
	RuleSetID       int32     `json:"ruleSetID" gorm:"column:ruleSetID"`
	Name            string    `json:"name"`
	RuleType        string    `json:"ruleType" gorm:"column:ruleType"`
	Statement       string    `json:"statement"`
	Actions         string    `json:"actions"`
	Status          int       `json:"status"`
	StatusCheck     int       `json:"statusCheck" gorm:"column:statusCheck"`
	StatusCheckText string    `json:"statusCheckText" gorm:"column:statusCheckText"`
	StatusCheckTime time.Time `json:"statusTime" gorm:"column:statusCheckTime"`
	CreateTime      time.Time `json:"createTime" gorm:"column:createTime"`
	UpdateTime      time.Time `json:"updateTime" gorm:"column:updateTime"`
}

func (Rule) TableName() string {
	return "rule"
}
