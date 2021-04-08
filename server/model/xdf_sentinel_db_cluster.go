// 自动生成模板SentinelDBClusterInfo
package model

import "strconv"

// 如果含有time.Time 请自行import time包

type SentinelDBClusterInfo struct {
	ID				  int    `json:"id" form:"id" gorm:"column:id; primarykey;AUTO_INCREMENT;comment:"`
	ClusterID         int    `json:"cluster_id" form:"cluster_id" gorm:"column:cluster_id;comment:"`
	LogicClusterID    int    `json:"logic_cluster_id" form:"logic_cluster_id" gorm:"column:logic_cluster_id;comment:"`
	LeaderEpoch       int    `json:"leader_epoch" form:"leader_epoch" gorm:"column:leader_epoch;comment:"`
	RlpcUser          string `json:"rlpc_user" form:"rlpc_user" gorm:"column:rlpc_user;comment:;type:char(30);size:30;"`
	RlpcPWD           string `json:"rlpc_pwd" form:"rlpc_pwd" gorm:"column:rlpc_pwd;comment:;type:char(30);size:30;"`
	SentinelClusterID int    `json:"sentinel_cluster_id" form:"sentinel_cluster_id" gorm:"column:sentinel_cluster_id;comment:"`
	Dbs               string `json:"dbs" form:"dbs" gorm:"column:dbs;type:varchar(1024) size:1024;comment:"`
}

func (info *SentinelDBClusterInfo)GetDBClusterName() string {
	return strconv.Itoa(info.ID)
}

func (SentinelDBClusterInfo) TableName() string {
	return "sentinel_db_cluster_info"
}

// 如果使用工作流功能 需要打开下方注释 并到initialize的workflow中进行注册 且必须指定TableName
// type SentinelDBClusterInfoWorkflow struct {
// 	// 工作流操作结构体
// 	WorkflowBase      `json:"wf"`
// 	SentinelDBClusterInfo   `json:"business"`
// }

// func (SentinelDBClusterInfo) TableName() string {
// 	return "sentinel_db_cluster_info"
// }

// 工作流注册代码

// initWorkflowModel内部注册
// model.WorkflowBusinessStruct["SentinelDBCluster"] = func() model.GVA_Workflow {
//   return new(model.SentinelDBClusterInfoWorkflow)
// }

// initWorkflowTable内部注册
// model.WorkflowBusinessTable["SentinelDBCluster"] = func() interface{} {
// 	return new(model.SentinelDBClusterInfo)
// }
