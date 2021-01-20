package service

import (
	"errors"
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
)


func GetSentinelClusterList() (err error, list interface{}, total int64) {
	// 创建db
	db := global.GVA_DB.Model(&model.SentinelClusterInfo{})
	var sentinelClusters []model.SentinelClusterInfo
	err = db.Count(&total).Error
	err = db.Find(&sentinelClusters).Error
	return err, sentinelClusters, total
}

//@description: 创建Tag记录
//@param: tag model.Tag
//@return: err error

func CreateSentinelClusterInfo(sentinelCluster model.SentinelClusterInfo) (id int, err error) {
	if global.GVA_DB.Where("name = ?", sentinelCluster.Name).Find(&[]model.SentinelClusterInfo{}).RowsAffected != 0 {
		return 0, errors.New(fmt.Sprintf( "sentinelCluster %s is existed", sentinelCluster.Name))
	}

	err = global.GVA_DB.Create(&sentinelCluster).Error

	return sentinelCluster.SentinelClusterID, err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteTag
//@description: 删除Tag记录
//@param: tag model.Tag
//@return: err error

func DeleteSentinelClusterInfo(sentinelCluster model.SentinelClusterInfo) (err error) {
	err = global.GVA_DB.Delete(&sentinelCluster).Error
	return err
}

