package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateBackUpDB
//@description: 创建BackUpDB记录
//@param: BackDB model.BackUpDB
//@return: err error

func CreateBackUpDB(BackDB model.BackUpDB) (err error) {
	err = global.GVA_DB.Create(&BackDB).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteBackUpDB
//@description: 删除BackUpDB记录
//@param: BackDB model.BackUpDB
//@return: err error

func DeleteBackUpDB(BackDB model.BackUpDB) (err error) {
	err = global.GVA_DB.Delete(&BackDB).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteBackUpDBByIds
//@description: 批量删除BackUpDB记录
//@param: ids request.IdsReq
//@return: err error

func DeleteBackUpDBByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.BackUpDB{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateBackUpDB
//@description: 更新BackUpDB记录
//@param: BackDB *model.BackUpDB
//@return: err error

func UpdateBackUpDB(BackDB model.BackUpDB) (err error) {
	err = global.GVA_DB.Save(&BackDB).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetBackUpDB
//@description: 根据id获取BackUpDB记录
//@param: id uint
//@return: err error, BackDB model.BackUpDB

func GetBackUpDB(id int) (err error, BackDB model.BackUpDB) {
	err = global.GVA_DB.Where("backup_id = ?", id).First(&BackDB).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetBackUpDBInfoList
//@description: 分页获取BackUpDB记录
//@param: info request.BackUpDBSearch
//@return: err error, list interface{}, total int64

func GetBackUpDBInfoList(info request.BackUpDBSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.BackUpDB{})
    var BackDBs []model.BackUpDB
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.IP != "" {
        db = db.Where("`ip` LIKE ?","%"+ info.IP+"%")
    }
    if info.Type != 0 {
        db = db.Where("`type` = ?",info.Type)
    }
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&BackDBs).Error
	return err, BackDBs, total
}

func GetAllBackUpDBs() (err error, list interface{}, total int64) {
	// 创建db
	db := global.GVA_DB.Model(&model.BackUpDB{})
	var BackDBs []model.BackUpDB
	err = db.Count(&total).Error
	err = db.Find(&BackDBs).Error
	return err, BackDBs, total
}
