package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateNode
//@description: 创建Node记录
//@param: node model.Node
//@return: err error

func CreateNode(node model.Node) (err error) {
	err = global.GVA_DB.Create(&node).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteNode
//@description: 删除Node记录
//@param: node model.Node
//@return: err error

func DeleteNode(node model.Node) (err error) {
	err = global.GVA_DB.Delete(&node).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteNodeByIds
//@description: 批量删除Node记录
//@param: ids request.IdsReq
//@return: err error

func DeleteNodeByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Node{},"node_id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateNode
//@description: 更新Node记录
//@param: node *model.Node
//@return: err error

func UpdateNode(node model.Node) (err error) {
	err = global.GVA_DB.Save(&node).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetNode
//@description: 根据id获取Node记录
//@param: id uint
//@return: err error, node model.Node

func GetNode(id int) (err error, node model.Node) {
	err = global.GVA_DB.Where("node_id = ?", id).First(&node).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetNodeInfoList
//@description: 分页获取Node记录
//@param: info request.NodeSearch
//@return: err error, list interface{}, total int64

func GetNodeInfoList(info request.NodeSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.Node{})
    var nodes []model.Node
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.TagID != 0 {
        db = db.Where("`tag_id` = ?",info.TagID)
    }
    if info.IP != "" {
        db = db.Where("`ip` LIKE ?","%"+ info.IP+"%")
    }
    if info.RoleID != 0 {
        db = db.Where("`role_id` = ?",info.RoleID)
    }
    if info.AreaID != 0 {
        db = db.Where("`area_id` = ?",info.AreaID)
    }
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&nodes).Error
	return err, nodes, total
}
