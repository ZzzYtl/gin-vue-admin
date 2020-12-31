// 自动生成模板DeptInfo
package model

// 如果含有time.Time 请自行import time包
type DeptInfo struct {
      DeptID  int `json:"dept_id" form:"dept_id" gorm:"primarykey;AUTO_INCREMENT;column:dept_id;comment:;type:int;size:10;"`
      DeptName  string `json:"dept_name" form:"dept_name" gorm:"column:dept_name;not null;comment:;type:char(30);size:30;"`
      LeaderName  string `json:"leader_name" form:"leader_name" gorm:"column:leader_name;not null;comment:;type:char(30);size:30;"`
}


func (DeptInfo) TableName() string {
  return "dept_info"
}


// 如果使用工作流功能 需要打开下方注释 并到initialize的workflow中进行注册 且必须指定TableName
// type DeptInfoWorkflow struct {
// 	// 工作流操作结构体
// 	WorkflowBase      `json:"wf"`
// 	DeptInfo   `json:"business"`
// }

// func (DeptInfo) TableName() string {
// 	return "dept_info"
// }

// 工作流注册代码

// initWorkflowModel内部注册
// model.WorkflowBusinessStruct["dept"] = func() model.GVA_Workflow {
//   return new(model.DeptInfoWorkflow)
// }

// initWorkflowTable内部注册
// model.WorkflowBusinessTable["dept"] = func() interface{} {
// 	return new(model.DeptInfo)
// }
