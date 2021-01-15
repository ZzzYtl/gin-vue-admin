// 自动生成模板Node
package model

// 如果含有time.Time 请自行import time包
type Node struct {
      NodeID  int `json:"node_id" form:"node_id" gorm:"primarykey;AUTO_INCREMENT;column:node_id;comment:;type:int;size:10;"`
      TagID  int `json:"tag_id" form:"tag_id" gorm:"column:tag_id;comment:;type:int;size:10;"`
      IP  string `json:"ip" form:"ip" gorm:"column:ip;comment:;type:char(15);size:15;"`
      RoleID  int `json:"role_id" form:"role_id" gorm:"column:role_id;comment:;type:smallint;size:3;"`
      AreaID  int `json:"area_id" form:"area_id" gorm:"column:area_id;comment:;type:smallint;size:3;"`
}


func (Node) TableName() string {
  return "node_info"
}


// 如果使用工作流功能 需要打开下方注释 并到initialize的workflow中进行注册 且必须指定TableName
// type NodeWorkflow struct {
// 	// 工作流操作结构体
// 	WorkflowBase      `json:"wf"`
// 	Node   `json:"business"`
// }

// func (Node) TableName() string {
// 	return "node_info"
// }

// 工作流注册代码

// initWorkflowModel内部注册
// model.WorkflowBusinessStruct["node"] = func() model.GVA_Workflow {
//   return new(model.NodeWorkflow)
// }

// initWorkflowTable内部注册
// model.WorkflowBusinessTable["node"] = func() interface{} {
// 	return new(model.Node)
// }
