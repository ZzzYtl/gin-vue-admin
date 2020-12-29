// 自动生成模板Area
package model

// 如果含有time.Time 请自行import time包
type Area struct {
      AreaID  int `json:"area_id" form:"area_id" gorm:"primarykey;AUTO_INCREMENT;column:area_id;comment:机房ID"`
      Name  string `json:"name" form:"name" gorm:"not null;column:name;comment:机房名称;type:char(30);size:30;"`
      Status  int `json:"status" form:"status" gorm:"column:status;comment:状态;type:smallint;size:4;"`
}


func (Area) TableName() string {
  return "area_info"
}


// 如果使用工作流功能 需要打开下方注释 并到initialize的workflow中进行注册 且必须指定TableName
// type AreaWorkflow struct {
// 	// 工作流操作结构体
// 	WorkflowBase      `json:"wf"`
// 	Area   `json:"business"`
// }

// func (Area) TableName() string {
// 	return "area_info"
// }

// 工作流注册代码

// initWorkflowModel内部注册
// model.WorkflowBusinessStruct["area"] = func() model.GVA_Workflow {
//   return new(model.AreaWorkflow)
// }

// initWorkflowTable内部注册
// model.WorkflowBusinessTable["area"] = func() interface{} {
// 	return new(model.Area)
// }
