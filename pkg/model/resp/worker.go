package resp

import (
	"bifromq_engine/pkg/model/entity"
	"time"
)

type WorkerRegisterResp struct {
	WorkerID int32 `json:"workerID"`
}

type WorkerListItem struct {
	entity.Worker
	StatusText string `json:"statusText"`
}

type WorkerListResp struct {
	PageData []WorkerListItem `json:"pageData"`
	Total    int64            `json:"total"`
}

type ConfigurationResp struct {
	Data           map[string]string `json:"data"`
	LastUpdateTime time.Time         `json:"lastUpdateTime"`
}

type WorkerHeartbeatResp struct {
	LastSourcesTime time.Time `json:"lastSourcesTime"`
	LastSinksTime   time.Time `json:"lastSinksTime"`
	Kill            bool      `json:"kill"`
}
