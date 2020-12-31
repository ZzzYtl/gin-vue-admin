import service from '@/utils/request'

// @Tags BackUpDB
// @Summary 创建BackUpDB
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BackUpDB true "创建BackUpDB"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /BackDB/createBackUpDB [post]
export const createBackUpDB = (data) => {
     return service({
         url: "/BackDB/createBackUpDB",
         method: 'post',
         data
     })
 }


// @Tags BackUpDB
// @Summary 删除BackUpDB
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BackUpDB true "删除BackUpDB"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /BackDB/deleteBackUpDB [delete]
 export const deleteBackUpDB = (data) => {
     return service({
         url: "/BackDB/deleteBackUpDB",
         method: 'delete',
         data
     })
 }

// @Tags BackUpDB
// @Summary 删除BackUpDB
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除BackUpDB"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /BackDB/deleteBackUpDB [delete]
 export const deleteBackUpDBByIds = (data) => {
     return service({
         url: "/BackDB/deleteBackUpDBByIds",
         method: 'delete',
         data
     })
 }

// @Tags BackUpDB
// @Summary 更新BackUpDB
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BackUpDB true "更新BackUpDB"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /BackDB/updateBackUpDB [put]
 export const updateBackUpDB = (data) => {
     return service({
         url: "/BackDB/updateBackUpDB",
         method: 'put',
         data
     })
 }


// @Tags BackUpDB
// @Summary 用id查询BackUpDB
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BackUpDB true "用id查询BackUpDB"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /BackDB/findBackUpDB [get]
 export const findBackUpDB = (params) => {
     return service({
         url: "/BackDB/findBackUpDB",
         method: 'get',
         params
     })
 }


// @Tags BackUpDB
// @Summary 分页获取BackUpDB列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取BackUpDB列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /BackDB/getBackUpDBList [get]
 export const getBackUpDBList = (params) => {
     return service({
         url: "/BackDB/getBackUpDBList",
         method: 'get',
         params
     })
 }