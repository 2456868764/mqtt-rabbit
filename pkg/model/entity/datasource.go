package entity

import "time"

type SourcesType string

const (
	SourcesMqtt  SourcesType = "mqtt"
	SourcesRedis SourcesType = "redis"
	SourcesKafka SourcesType = "kafka"
)

type ConfType string

const (
	ConfTypeGlobal ConfType = "global"
	ConfTypeSource ConfType = "source"
	ConfTypeSink   ConfType = "sink"
)

type DataSource struct {
	ID         int
	Name       string
	ConfType   string `json:"confType" gorm:"column:confType"`
	Type       string
	Content    string
	CreateTime time.Time `json:"createTime" gorm:"column:createTime"`
	UpdateTime time.Time `json:"updateTime" gorm:"column:updateTime"`
}

func (DataSource) TableName() string {
	return "data_source"
}

type Schema struct {
	ID         int
	Name       string
	Statement  string
	CreateTime time.Time `json:"createTime" gorm:"column:createTime"`
	UpdateTime time.Time `json:"updateTime" gorm:"column:updateTime"`
}

func (Schema) TableName() string {
	return "schema"
}
