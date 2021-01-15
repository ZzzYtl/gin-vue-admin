package request

import "gin-vue-admin/model"

type DataBaseInfoSearch struct{
    model.DataBaseInfo
    PageInfo
}