package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateSentinelInfo
//@description: 创建SentinelInfo记录
//@param: sentinelinfo model.SentinelInfo
//@return: err error

func CreateSentinelInfo(sentinelinfo model.SentinelInfo) (err error) {
	err = global.GVA_DB.Create(&sentinelinfo).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteSentinelInfo
//@description: 删除SentinelInfo记录
//@param: sentinelinfo model.SentinelInfo
//@return: err error

func DeleteSentinelInfo(sentinelinfo model.SentinelInfo) (err error) {
	err = global.GVA_DB.Delete(&sentinelinfo).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteSentinelInfoByIds
//@description: 批量删除SentinelInfo记录
//@param: ids request.IdsReq
//@return: err error

func DeleteSentinelInfoByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.SentinelInfo{},"sentinel_id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateSentinelInfo
//@description: 更新SentinelInfo记录
//@param: sentinelinfo *model.SentinelInfo
//@return: err error

func UpdateSentinelInfo(sentinelinfo model.SentinelInfo) (err error) {
	err = global.GVA_DB.Save(&sentinelinfo).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetSentinelInfo
//@description: 根据id获取SentinelInfo记录
//@param: id uint
//@return: err error, sentinelinfo model.SentinelInfo

func GetSentinelInfo(id int) (err error, sentinelinfo model.SentinelInfo) {
	err = global.GVA_DB.Where("sentinel_id = ?", id).First(&sentinelinfo).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetSentinelInfoInfoList
//@description: 分页获取SentinelInfo记录
//@param: info request.SentinelInfoSearch
//@return: err error, list interface{}, total int64

func GetSentinelInfoInfoList(info request.SentinelInfoSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.SentinelInfo{})
    var sentinelinfos []model.SentinelInfo
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.SentinelClusterID != 0 {
        db = db.Where("`sentinel_cluster_id` = ?",info.SentinelClusterID)
    }
    if info.IP != "" {
        db = db.Where("`ip` LIKE ?","%"+ info.IP+"%")
    }
    if info.Port != 0 {
        db = db.Where("`port` = ?",info.Port)
    }
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&sentinelinfos).Error
	return err, sentinelinfos, total
}
