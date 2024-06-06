package entity

import "time"

type DataSourceType string

const (
	DataSourceGlobal DataSourceType = "global"
	DataSourceMqtt   DataSourceType = "mqtt"
	DataSourceRedis  DataSourceType = "redis"
	DataSourceKafka  DataSourceType = "kafka"
)

type DataSource struct {
	ID         int
	Name       string
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
