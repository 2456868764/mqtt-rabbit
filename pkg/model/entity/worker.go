package entity

import "time"

type WorkerStatus int32

const (
	WorkerStatusOK            WorkerStatus = 1
	WorkerStatusHeartbeatMiss WorkerStatus = 2
	WorkerStatusRegister      WorkerStatus = 3
	WorkerStatusUnRegister    WorkerStatus = 4
	WorkerStatusDeleted       WorkerStatus = 5
)

// Enum value maps for WorkerStatus.
var (
	WorkerStatus_name = map[int32]string{
		0: "OK",
		1: "HeartbeatMiss",
		2: "Register",
		3: "UnRegister",
	}
	WorkerStatus_value = map[string]int32{
		"OK":            0,
		"HeartbeatMiss": 1,
		"Register":      2,
		"UnRegister":    3,
	}
)

type Worker struct {
	ID              int32     `json:"id"`
	Name            string    `json:"name"`
	Tag             string    `json:"tag"`
	IP              string    `json:"ip"`
	Status          int32     `json:"status" `
	Port            int32     `json:"port"`
	HeartbeatMisses int32     `json:"heartbeatMisses" gorm:"column:heartbeatMisses"`
	HeartbeatTime   time.Time `json:"heartbeatTime" gorm:"column:heartbeatTime;default:NULL"`
	LastSourcesTime time.Time `json:"lastSourcesTime" gorm:"column:lastSourcesTime;default:NULL"`
	CreateTime      time.Time `json:"createTime" gorm:"column:createTime"`
	UpdateTime      time.Time `json:"updateTime" gorm:"column:updateTime"`
}

func (Worker) TableName() string {
	return "worker"
}
