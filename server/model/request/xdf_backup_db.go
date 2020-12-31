package request

import "gin-vue-admin/model"

type BackUpDBSearch struct{
    model.BackUpDB
    PageInfo
}