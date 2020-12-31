package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateArea
//@description: 创建Area记录
//@param: area model.Area
//@return: err error

func CreateArea(area model.Area) (err error) {
	err = global.GVA_DB.Create(&area).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteArea
//@description: 删除Area记录
//@param: area model.Area
//@return: err error

func DeleteArea(area model.Area) (err error) {
	err = global.GVA_DB.Delete(&area).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteAreaByIds
//@description: 批量删除Area记录
//@param: ids request.IdsReq
//@return: err error

func DeleteAreaByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Area{},"area_id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateArea
//@description: 更新Area记录
//@param: area *model.Area
//@return: err error

func UpdateArea(area model.Area) (err error) {
	err = global.GVA_DB.Save(&area).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetArea
//@description: 根据id获取Area记录
//@param: id uint
//@return: err error, area model.Area

func GetArea(id int) (err error, area model.Area) {
	err = global.GVA_DB.Where("area_id = ?", id).First(&area).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetAreaInfoList
//@description: 分页获取Area记录
//@param: info request.AreaSearch
//@return: err error, list interface{}, total int64

func GetAreaInfoList(info request.AreaSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.Area{})
    var areas []model.Area
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Name != "" {
        db = db.Where("`name` LIKE ?","%"+ info.Name+"%")
    }
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&areas).Error
	return err, areas, total
}
