package request

import "gin-vue-admin/model"

type AreaSearch struct{
    model.Area
    PageInfo
}