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
	WorkerID           int32     `form:"workerId" binding:"required"`
	IP                 string    `form:"ip" binding:"required"`
	Port               int32     `form:"port" binding:"required"`
	LastDatasourceTime time.Time `form:"lastDataSourceTime"`
}

type ConfigurationReq struct {
	WorkerID int32  `form:"workerId" binding:"required"`
	Type     string `form:"type" binding:"required"`
}
