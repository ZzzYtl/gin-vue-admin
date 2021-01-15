package request

import "gin-vue-admin/model"

type NodeSearch struct{
    model.Node
    PageInfo
}