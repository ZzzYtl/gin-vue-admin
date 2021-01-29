package service

import (
	"errors"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateDataBaseInfo
//@description: 创建DataBaseInfo记录
//@param: DBInfo model.DataBaseInfo
//@return: err error

func CreateDataBaseInfo(DBInfo model.DataBaseInfo) (err error) {
	err = global.GVA_DB.Create(&DBInfo).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteDataBaseInfo
//@description: 删除DataBaseInfo记录
//@param: DBInfo model.DataBaseInfo
//@return: err error

func DeleteDataBaseInfo(DBInfo model.DataBaseInfo) (err error) {
	db := global.GVA_DB.Model(&model.SentinelDBClusterInfo{})
	var total int64
	db = db.Where("`cluster_id` = ?",DBInfo.ClusterID)
	err = db.Count(&total).Error
	if err != nil {
		return err
	}
	if total != 0 {
		return errors.New("cant delete, this db cluster is under sentinel")
	}
	err = global.GVA_DB.Delete(&DBInfo).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteDataBaseInfoByIds
//@description: 批量删除DataBaseInfo记录
//@param: ids request.IdsReq
//@return: err error

func DeleteDataBaseInfoByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.DataBaseInfo{},"cluster_id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateDataBaseInfo
//@description: 更新DataBaseInfo记录
//@param: DBInfo *model.DataBaseInfo
//@return: err error

func UpdateDataBaseInfo(DBInfo model.DataBaseInfo) (err error) {
	err = global.GVA_DB.Save(&DBInfo).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetDataBaseInfo
//@description: 根据id获取DataBaseInfo记录
//@param: id uint
//@return: err error, DBInfo model.DataBaseInfo

func GetDataBaseInfo(id int) (err error, DBInfo model.DataBaseInfo) {
	err = global.GVA_DB.Where("cluster_id = ?", id).First(&DBInfo).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetDataBaseInfoInfoList
//@description: 分页获取DataBaseInfo记录
//@param: info request.DataBaseInfoSearch
//@return: err error, list interface{}, total int64

func GetDataBaseInfoInfoList(info request.DataBaseInfoSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.DataBaseInfo{})
    var DBInfos []model.DataBaseInfo
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.TagID != 0 {
        db = db.Where("`tag_id` = ?",info.TagID)
    }
    if info.ClusterName != "" {
        db = db.Where("`cluster_name` LIKE ?","%"+ info.ClusterName+"%")
    }
    if info.DBType != 0 {
        db = db.Where("`cluster_type` = ?",info.DBType)
    }
    if info.DeptID != 0 {
        db = db.Where("`dept_id` = ?",info.DeptID)
    }
    if info.Port != 0 {
        db = db.Where("`port` = ?",info.Port)
    }
    if info.DBName != "" {
        db = db.Where("`db_name` LIKE ?","%"+ info.DBName+"%")
    }
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&DBInfos).Error
	return err, DBInfos, total
}

func GetDBList() (err error, list interface{}, total int64) {
	// 创建db
	db := global.GVA_DB.Model(&model.DataBaseInfo{})
	var dbs []model.DataBaseInfo
	err = db.Count(&total).Error
	err = db.Find(&dbs).Error
	return err, dbs, total
}
