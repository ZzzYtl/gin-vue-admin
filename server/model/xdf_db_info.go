// 自动生成模板DataBaseInfo
package model

// 如果含有time.Time 请自行import time包
type DataBaseInfo struct {
      ClusterID  int `json:"cluster_id" form:"cluster_id" gorm:"column:cluster_id;primarykey; AUTO_INCREMENT;comment:ID;type:int;size:10;"`
      TagID  int `json:"tag_id" form:"tag_id" gorm:"column:tag_id;comment:;type:int;size:10;"`
      ClusterName  string `json:"cluster_name" form:"cluster_name" gorm:"column:cluster_name;comment:;type:char(30);size:30;"`
      DBType  int `json:"db_type" form:"db_type" gorm:"column:cluster_type;comment:;type:smallint;size:3;"`
      DeptID  int `json:"dept_id" form:"dept_id" gorm:"column:dept_id;comment:;type:int;size:10;"`
      Port  int `json:"port" form:"port" gorm:"column:port;comment:;type:int;size:10;"`
      DBName  string `json:"db_name" form:"db_name" gorm:"column:db_name;comment:;type:char(30);size:30;"`
      DBUser  string `json:"db_user" form:"db_user" gorm:"column:db_user;comment:;type:char(30);size:30;"`
      DBPWD  string `json:"db_pwd" form:"db_pwd" gorm:"column:db_pwd;comment:;type:char(30);size:30;"`
      ZKName  string `json:"zk_name" form:"zk_name" gorm:"column:zk_name;comment:;type:char(30);size:30;"`
      BackUpID  int `json:"backup_id" form:"backup_id" gorm:"column:backup_id;comment:;type:smallint;size:3;"`
      DelayBackUp  int `json:"delay_backup_id" form:"delay_backup_id" gorm:"column:delay_backup_id;comment:;type:smallint;size:3;"`
}


func (DataBaseInfo) TableName() string {
  return "cluster_info"
}


// 如果使用工作流功能 需要打开下方注释 并到initialize的workflow中进行注册 且必须指定TableName
// type DataBaseInfoWorkflow struct {
// 	// 工作流操作结构体
// 	WorkflowBase      `json:"wf"`
// 	DataBaseInfo   `json:"business"`
// }

// func (DataBaseInfo) TableName() string {
// 	return "cluster_info"
// }

// 工作流注册代码

// initWorkflowModel内部注册
// model.WorkflowBusinessStruct["DBInfo"] = func() model.GVA_Workflow {
//   return new(model.DataBaseInfoWorkflow)
// }

// initWorkflowTable内部注册
// model.WorkflowBusinessTable["DBInfo"] = func() interface{} {
// 	return new(model.DataBaseInfo)
// }
