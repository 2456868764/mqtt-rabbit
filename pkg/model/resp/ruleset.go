package resp

import (
	"bifromq_engine/pkg/model/entity"
)

type RuleSetListItem struct {
	entity.RuleSet
	Node            *RuleSetNode `json:"node"`
	StreamCount     int          `json:"streamCount"`
	RuleCount       int          `json:"ruleCount"`
	StatusText      string       `json:"statusText"`
	StatusCheckText string       `json:"statusCheckText"`
}

type RuleSetNode struct {
	ID     int32  `json:"id"`
	Name   string `json:"name"`
	Tag    string `json:"tag"`
	IP     string `json:"ip"`
	Status int32  `json:"status" `
	Port   int32  `json:"port"`
}

type RuleSetListResp struct {
	PageData []RuleSetListItem `json:"pageData"`
	Total    int64             `json:"total"`
}

type RuleSetResp struct {
	entity.RuleSet
	StatusCheckText string            `json:"statusCheckText"`
	Streams         []*StreamListItem `json:"streams"`
	Rules           []*RuleListItem   `json:"rules"`
	Node            *RuleSetNode      `json:"node"`
}

type StreamListItem struct {
	entity.Stream
	StatusText string `json:"statusText"`
}

type RuleListItem struct {
	entity.Rule
	StatusText      string `json:"statusText"`
	StatusCheckDesc string `json:"statusCheckDesc"`
}
