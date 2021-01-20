package request

import "gin-vue-admin/model"

type SentinelInfoSearch struct{
    model.SentinelInfo
    PageInfo
}