package request

import "gin-vue-admin/model"

type SentinelDBClusterInfoSearch struct{
    model.SentinelDBClusterInfo
    PageInfo
}