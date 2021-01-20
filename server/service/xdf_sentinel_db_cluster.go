package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateSentinelDBClusterInfo
//@description: 创建SentinelDBClusterInfo记录
//@param: SentinelDBCluster model.SentinelDBClusterInfo
//@return: err error

func CreateSentinelDBClusterInfo(SentinelDBCluster model.SentinelDBClusterInfo) (err error) {
	err = global.GVA_DB.Create(&SentinelDBCluster).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteSentinelDBClusterInfo
//@description: 删除SentinelDBClusterInfo记录
//@param: SentinelDBCluster model.SentinelDBClusterInfo
//@return: err error

func DeleteSentinelDBClusterInfo(SentinelDBCluster model.SentinelDBClusterInfo) (err error) {
	err = global.GVA_DB.Delete(&SentinelDBCluster).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteSentinelDBClusterInfoByIds
//@description: 批量删除SentinelDBClusterInfo记录
//@param: ids request.IdsReq
//@return: err error

func DeleteSentinelDBClusterInfoByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.SentinelDBClusterInfo{},"cluster_id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateSentinelDBClusterInfo
//@description: 更新SentinelDBClusterInfo记录
//@param: SentinelDBCluster *model.SentinelDBClusterInfo
//@return: err error

func UpdateSentinelDBClusterInfo(SentinelDBCluster model.SentinelDBClusterInfo) (err error) {
	err = global.GVA_DB.Save(&SentinelDBCluster).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetSentinelDBClusterInfo
//@description: 根据id获取SentinelDBClusterInfo记录
//@param: id uint
//@return: err error, SentinelDBCluster model.SentinelDBClusterInfo

func GetSentinelDBClusterInfo(id int) (err error, SentinelDBCluster model.SentinelDBClusterInfo) {
	err = global.GVA_DB.Where("cluster_id = ?", id).First(&SentinelDBCluster).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetSentinelDBClusterInfoInfoList
//@description: 分页获取SentinelDBClusterInfo记录
//@param: info request.SentinelDBClusterInfoSearch
//@return: err error, list interface{}, total int64

func GetSentinelDBClusterInfoInfoList(info request.SentinelDBClusterInfoSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.SentinelDBClusterInfo{})
    var SentinelDBClusters []model.SentinelDBClusterInfo
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.SentinelClusterID != 0 {
        db = db.Where("`sentinel_cluster_id` = ?",info.SentinelClusterID)
    }
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&SentinelDBClusters).Error
	return err, SentinelDBClusters, total
}
