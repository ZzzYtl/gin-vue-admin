package service

import (
	"errors"
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
)


func GetTagList() (err error, list interface{}, total int64) {
    // 创建db
	db := global.GVA_DB.Model(&model.Tag{})
    var tags []model.Tag
	err = db.Count(&total).Error
	err = db.Find(&tags).Error
	return err, tags, total
}

//@description: 创建Tag记录
//@param: tag model.Tag
//@return: err error

func CreateTag(tag model.Tag) (id int, err error) {
	if global.GVA_DB.Where("name = ?", tag.Name).Find(&[]model.Tag{}).RowsAffected != 0 {
		return 0, errors.New(fmt.Sprintf( "tag %s is existed", tag.Name))
	}

	err = global.GVA_DB.Create(&tag).Error

	return tag.TagID, err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteTag
//@description: 删除Tag记录
//@param: tag model.Tag
//@return: err error

func DeleteTag(tag model.Tag) (err error) {
	err = global.GVA_DB.Delete(&tag).Error
	return err
}
