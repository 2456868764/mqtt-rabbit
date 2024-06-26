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
		0: "INIT",
		1: "QUEUED",
		2: "STARTED",
		3: "COMPLETE",
		4: "STOPPED",
		5: "FAILED",
	}
	TaskStatus_value = map[string]int32{
		"INIT":     0,
		"QUEUED":   1,
		"STARTED":  2,
		"COMPLETE": 3,
		"STOPPED":  4,
		"FAILED":   5,
	}
)

var (
	TaskCheckStatus_name = map[int32]string{
		0: "INIT",
		1: "RUNNING",
		2: "STOPPED",
		3: "FAILED",
		4: "PART FAILED",
	}
	TaskCheckStatus_value = map[string]int32{
		"INIT":        0,
		"RUNNING":     1,
		"STOPPED":     2,
		"FAILED":      3,
		"PART FAILED": 4,
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
	Status          int32     `json:"status"`
	Deleted         int       `json:"deleted"`
	WorkerID        int32     `json:"WorkerID"  gorm:"column:workerID"`
	StatusCheck     int32     `json:"statusCheck" gorm:"column:statusCheck"`
	StatusCheckTime time.Time `json:"statusCheckTime" gorm:"column:statusCheckTime;default:NULL"`
	ScheduleTime    time.Time `json:"scheduleTime" gorm:"column:scheduleTime;default:NULL"`
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
	Status     int32     `json:"status"`
	Deleted    int       `json:"deleted"`
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
	Status          int32     `json:"status"`
	Deleted         int       `json:"deleted"`
	StatusCheck     int32     `json:"statusCheck" gorm:"column:statusCheck"`
	StatusCheckText string    `json:"statusCheckText" gorm:"column:statusCheckText"`
	StatusCheckTime time.Time `json:"statusTime" gorm:"column:statusCheckTime"`
	CreateTime      time.Time `json:"createTime" gorm:"column:createTime"`
	UpdateTime      time.Time `json:"updateTime" gorm:"column:updateTime"`
}

func (Rule) TableName() string {
	return "rule"
}
