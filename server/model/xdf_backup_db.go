// 自动生成模板BackUpDB
package model

// 如果含有time.Time 请自行import time包
type BackUpDB struct {
      BackUpID  int `json:"backup_id" form:"backup_id" gorm:"primarykey;AUTO_INCREMENT;column:backup_id;comment:;type:int;"`
      IP  string `json:"IP" form:"IP" gorm:"column:ip;not null;comment:ip;type:char(15);size:15;"`
      Type  int `json:"type" form:"type" gorm:"column:type;comment:类型;type:smallint;size:3;"`
      Domain string `json:"domain" form:"domain" gorm:"column:domain;not null;comment:domain;type:char(128);size:128;"`
}


func (BackUpDB) TableName() string {
  return "backup_info"
}


// 如果使用工作流功能 需要打开下方注释 并到initialize的workflow中进行注册 且必须指定TableName
// type BackUpDBWorkflow struct {
// 	// 工作流操作结构体
// 	WorkflowBase      `json:"wf"`
// 	BackUpDB   `json:"business"`
// }

// func (BackUpDB) TableName() string {
// 	return "backup_info"
// }

// 工作流注册代码

// initWorkflowModel内部注册
// model.WorkflowBusinessStruct["BackDB"] = func() model.GVA_Workflow {
//   return new(model.BackUpDBWorkflow)
// }

// initWorkflowTable内部注册
// model.WorkflowBusinessTable["BackDB"] = func() interface{} {
// 	return new(model.BackUpDB)
// }
