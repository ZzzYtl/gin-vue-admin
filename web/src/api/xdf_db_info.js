import service from '@/utils/request'

// @Tags DataBaseInfo
// @Summary 创建DataBaseInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DataBaseInfo true "创建DataBaseInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /DBInfo/createDataBaseInfo [post]
export const createDataBaseInfo = (data) => {
     return service({
         url: "/DBInfo/createDataBaseInfo",
         method: 'post',
         data
     })
 }


// @Tags DataBaseInfo
// @Summary 删除DataBaseInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DataBaseInfo true "删除DataBaseInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /DBInfo/deleteDataBaseInfo [delete]
 export const deleteDataBaseInfo = (data) => {
     return service({
         url: "/DBInfo/deleteDataBaseInfo",
         method: 'delete',
         data
     })
 }

// @Tags DataBaseInfo
// @Summary 删除DataBaseInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除DataBaseInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /DBInfo/deleteDataBaseInfo [delete]
 export const deleteDataBaseInfoByIds = (data) => {
     return service({
         url: "/DBInfo/deleteDataBaseInfoByIds",
         method: 'delete',
         data
     })
 }

// @Tags DataBaseInfo
// @Summary 更新DataBaseInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DataBaseInfo true "更新DataBaseInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /DBInfo/updateDataBaseInfo [put]
 export const updateDataBaseInfo = (data) => {
     return service({
         url: "/DBInfo/updateDataBaseInfo",
         method: 'put',
         data
     })
 }


// @Tags DataBaseInfo
// @Summary 用id查询DataBaseInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DataBaseInfo true "用id查询DataBaseInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /DBInfo/findDataBaseInfo [get]
 export const findDataBaseInfo = (params) => {
     return service({
         url: "/DBInfo/findDataBaseInfo",
         method: 'get',
         params
     })
 }


// @Tags DataBaseInfo
// @Summary 分页获取DataBaseInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取DataBaseInfo列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /DBInfo/getDataBaseInfoList [get]
 export const getDataBaseInfoList = (params) => {
     return service({
         url: "/DBInfo/getDataBaseInfoList",
         method: 'get',
         params
     })
 }

 export const getAllDBs = (data) => {
    return service({
        url: "/DBInfo/getAllDBs",
        method: 'post',
        data
    })
}