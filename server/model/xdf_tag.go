// 自动生成模板Node
package model

// 如果含有time.Time 请自行import time包
type Tag struct {
	TagID  int `json:"tag_id" form:"tag_id" gorm:"primarykey;AUTO_INCREMENT;column:tag_id;comment:;type:int;size:10;"`
	Name  string `json:"name" form:"name" gorm:"column:name; unique; comment:;type:char(30);size:30;"`
}


func (Tag) TableName() string {
	return "tag_info"
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
