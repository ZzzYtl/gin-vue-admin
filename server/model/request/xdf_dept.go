package request

import "gin-vue-admin/model"

type DeptInfoSearch struct{
    model.DeptInfo
    PageInfo
}