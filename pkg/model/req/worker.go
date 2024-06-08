package req

import "time"

type WorkerRegisterReq struct {
	Name string `form:"name" binding:"required"`
	Tag  string `form:"tag"`
	IP   string `form:"ip" binding:"required"`
	Port int32  `form:"port" binding:"required"`
}

type WorkerUnRegisterReq struct {
	WorkerID int32 `form:"workerId" binding:"required"`
}

type WorkerHeartbeatReq struct {
	WorkerID        int32     `form:"workerId" binding:"required"`
	IP              string    `form:"ip" binding:"required"`
	Port            int32     `form:"port" binding:"required"`
	LastSourcesTime time.Time `form:"lastSourcesTime"`
}

type ConfigurationReq struct {
	WorkerID int32  `form:"workerId" binding:"required"`
	ConfType string `form:"confType" binding:"required"`
}
