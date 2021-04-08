package service

import (
	"errors"
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
)


func GetLogicClusterList() (err error, list interface{}, total int64) {
	// 创建db
	db := global.GVA_DB.Model(&model.LogicCluster{})
	var lcs []model.LogicCluster
	err = db.Count(&total).Error
	err = db.Find(&lcs).Error
	return err, lcs, total
}

//@description: 创建LogicCluster记录
//@param: logiccluster model.LogicCluster
//@return: err error

func CreateLogicCluster(logicCluster model.LogicCluster) (id int, err error) {
	if global.GVA_DB.Where("name = ?", logicCluster.Name).Find(&[]model.LogicCluster{}).RowsAffected != 0 {
		return 0, errors.New(fmt.Sprintf( "logicCluster %s is existed", logicCluster.Name))
	}

	err = global.GVA_DB.Create(&logicCluster).Error

	return logicCluster.LogicClusterID, err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteLogicCluster
//@description: 删除LogicCluster记录
//@param: logicCluster model.LogicCluster
//@return: err error

func DeleteLogicCluster(lc model.LogicCluster) (err error) {
	err = global.GVA_DB.Delete(&lc).Error
	return err
}
