package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateDeptInfo
//@description: 创建DeptInfo记录
//@param: dept model.DeptInfo
//@return: err error

func CreateDeptInfo(dept model.DeptInfo) (err error) {
	err = global.GVA_DB.Create(&dept).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteDeptInfo
//@description: 删除DeptInfo记录
//@param: dept model.DeptInfo
//@return: err error

func DeleteDeptInfo(dept model.DeptInfo) (err error) {
	err = global.GVA_DB.Delete(&dept).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteDeptInfoByIds
//@description: 批量删除DeptInfo记录
//@param: ids request.IdsReq
//@return: err error

func DeleteDeptInfoByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.DeptInfo{},"dept_id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateDeptInfo
//@description: 更新DeptInfo记录
//@param: dept *model.DeptInfo
//@return: err error

func UpdateDeptInfo(dept model.DeptInfo) (err error) {
	err = global.GVA_DB.Save(&dept).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetDeptInfo
//@description: 根据id获取DeptInfo记录
//@param: id uint
//@return: err error, dept model.DeptInfo

func GetDeptInfo(id int) (err error, dept model.DeptInfo) {
	err = global.GVA_DB.Where("dept_id = ?", id).First(&dept).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetDeptInfoInfoList
//@description: 分页获取DeptInfo记录
//@param: info request.DeptInfoSearch
//@return: err error, list interface{}, total int64

func GetDeptInfoInfoList(info request.DeptInfoSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.DeptInfo{})
    var depts []model.DeptInfo
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.DeptName != "" {
        db = db.Where("`dept_name` LIKE ?","%"+ info.DeptName+"%")
    }
    if info.LeaderName != "" {
        db = db.Where("`leader_name` LIKE ?","%"+ info.LeaderName+"%")
    }
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&depts).Error
	return err, depts, total
}

func GetDeptList() (err error, list interface{}, total int64) {
	// 创建db
	db := global.GVA_DB.Model(&model.DeptInfo{})
	var depts []model.DeptInfo
	err = db.Count(&total).Error
	err = db.Find(&depts).Error
	return err, depts, total
}
