import service from '@/utils/request'

// @Tags SentinelDBClusterInfo
// @Summary 创建SentinelDBClusterInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SentinelDBClusterInfo true "创建SentinelDBClusterInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /SentinelDBCluster/createSentinelDBClusterInfo [post]
export const createSentinelDBClusterInfo = (data) => {
     return service({
         url: "/SentinelDBCluster/createSentinelDBClusterInfo",
         method: 'post',
         data
     })
 }


// @Tags SentinelDBClusterInfo
// @Summary 删除SentinelDBClusterInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SentinelDBClusterInfo true "删除SentinelDBClusterInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /SentinelDBCluster/deleteSentinelDBClusterInfo [delete]
 export const deleteSentinelDBClusterInfo = (data) => {
     return service({
         url: "/SentinelDBCluster/deleteSentinelDBClusterInfo",
         method: 'delete',
         data
     })
 }

 export const changeOldMaster2Slave = (data) => {
    return service({
        url: "/SentinelDBCluster/changeOldMaster2Slave",
        method: 'post',
        data
    })
}

// @Tags SentinelDBClusterInfo
// @Summary 删除SentinelDBClusterInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SentinelDBClusterInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /SentinelDBCluster/deleteSentinelDBClusterInfo [delete]
 export const deleteSentinelDBClusterInfoByIds = (data) => {
     return service({
         url: "/SentinelDBCluster/deleteSentinelDBClusterInfoByIds",
         method: 'delete',
         data
     })
 }

// @Tags SentinelDBClusterInfo
// @Summary 更新SentinelDBClusterInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SentinelDBClusterInfo true "更新SentinelDBClusterInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /SentinelDBCluster/updateSentinelDBClusterInfo [put]
 export const updateSentinelDBClusterInfo = (data) => {
     return service({
         url: "/SentinelDBCluster/updateSentinelDBClusterInfo",
         method: 'put',
         data
     })
 }


// @Tags SentinelDBClusterInfo
// @Summary 用id查询SentinelDBClusterInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SentinelDBClusterInfo true "用id查询SentinelDBClusterInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /SentinelDBCluster/findSentinelDBClusterInfo [get]
 export const findSentinelDBClusterInfo = (params) => {
     return service({
         url: "/SentinelDBCluster/findSentinelDBClusterInfo",
         method: 'get',
         params
     })
 }


// @Tags SentinelDBClusterInfo
// @Summary 分页获取SentinelDBClusterInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取SentinelDBClusterInfo列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /SentinelDBCluster/getSentinelDBClusterInfoList [get]
 export const getSentinelDBClusterInfoList = (params) => {
     return service({
         url: "/SentinelDBCluster/getSentinelDBClusterInfoList",
         method: 'get',
         params
     })
 }