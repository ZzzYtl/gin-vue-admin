import service from '@/utils/request'

// @Tags SentinelInfo
// @Summary 创建SentinelInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SentinelInfo true "创建SentinelInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sentinelinfo/createSentinelInfo [post]
export const createSentinelInfo = (data) => {
     return service({
         url: "/sentinelinfo/createSentinelInfo",
         method: 'post',
         data
     })
 }


// @Tags SentinelInfo
// @Summary 删除SentinelInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SentinelInfo true "删除SentinelInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sentinelinfo/deleteSentinelInfo [delete]
 export const deleteSentinelInfo = (data) => {
     return service({
         url: "/sentinelinfo/deleteSentinelInfo",
         method: 'delete',
         data
     })
 }

// @Tags SentinelInfo
// @Summary 删除SentinelInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SentinelInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sentinelinfo/deleteSentinelInfo [delete]
 export const deleteSentinelInfoByIds = (data) => {
     return service({
         url: "/sentinelinfo/deleteSentinelInfoByIds",
         method: 'delete',
         data
     })
 }

// @Tags SentinelInfo
// @Summary 更新SentinelInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SentinelInfo true "更新SentinelInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sentinelinfo/updateSentinelInfo [put]
 export const updateSentinelInfo = (data) => {
     return service({
         url: "/sentinelinfo/updateSentinelInfo",
         method: 'put',
         data
     })
 }


// @Tags SentinelInfo
// @Summary 用id查询SentinelInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SentinelInfo true "用id查询SentinelInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sentinelinfo/findSentinelInfo [get]
 export const findSentinelInfo = (params) => {
     return service({
         url: "/sentinelinfo/findSentinelInfo",
         method: 'get',
         params
     })
 }


// @Tags SentinelInfo
// @Summary 分页获取SentinelInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取SentinelInfo列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sentinelinfo/getSentinelInfoList [get]
 export const getSentinelInfoList = (params) => {
     return service({
         url: "/sentinelinfo/getSentinelInfoList",
         method: 'get',
         params
     })
 }