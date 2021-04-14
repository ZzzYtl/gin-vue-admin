// 自动生成模板SentinelInfo
package model


// 如果含有time.Time 请自行import time包
type SentinelInfo struct {
      SentinelID  int `json:"sentinel_id" form:"sentinel_id" gorm:"primarykey; AUTO_INCREMENT; column:sentinel_id;comment:"`
      RunID  string `json:"run_id" form:"run_id" gorm:"column:run_id;comment:;type:char(30);size:30;"`
      SentinelClusterID  int `json:"sentinel_cluster_id" form:"sentinel_cluster_id" gorm:"column:sentinel_cluster_id;comment:"`
      IP  string `json:"ip" form:"ip" gorm:"column:ip;comment:;type:char(15);size:15;"`
      Port  int `json:"port" form:"port" gorm:"column:port;comment:"`
      CurrentEpoch  int `json:"current_epoch" form:"current_epoch" gorm:"column:current_epoch;comment:"`
      ConfigEpoch  int `json:"config_epoch" form:"config_epoch" gorm:"column:config_epoch;comment:"`
      OnLine  bool `json:"online" form:"online"`
}


func (SentinelInfo) TableName() string {
  return "sentinel_info"
}


// 如果使用工作流功能 需要打开下方注释 并到initialize的workflow中进行注册 且必须指定TableName
// type SentinelInfoWorkflow struct {
// 	// 工作流操作结构体
// 	WorkflowBase      `json:"wf"`
// 	SentinelInfo   `json:"business"`
// }

// func (SentinelInfo) TableName() string {
// 	return "sentinel_info"
// }

// 工作流注册代码

// initWorkflowModel内部注册
// model.WorkflowBusinessStruct["sentinelinfo"] = func() model.GVA_Workflow {
//   return new(model.SentinelInfoWorkflow)
// }

// initWorkflowTable内部注册
// model.WorkflowBusinessTable["sentinelinfo"] = func() interface{} {
// 	return new(model.SentinelInfo)
// }
