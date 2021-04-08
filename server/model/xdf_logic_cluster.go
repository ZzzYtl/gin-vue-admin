// 自动生成模板Node
package model

// 如果含有time.Time 请自行import time包
type LogicCluster struct {
	LogicClusterID  int `json:"logic_cluster_id" form:"logic_cluster_id" gorm:"primarykey;AUTO_INCREMENT;column:logic_cluster_id;comment:;type:int;size:10;"`
	Name  string `json:"name" form:"name" gorm:"column:name; unique; comment:;type:char(30);size:30;"`
}


func (LogicCluster) TableName() string {
	return "logic_cluster"
}
