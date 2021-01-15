import service from '@/utils/request'

// @Tags Node
// @Summary 创建Node
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Node true "创建Node"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /node/createNode [post]
export const createNode = (data) => {
     return service({
         url: "/node/createNode",
         method: 'post',
         data
     })
 }


// @Tags Node
// @Summary 删除Node
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Node true "删除Node"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /node/deleteNode [delete]
 export const deleteNode = (data) => {
     return service({
         url: "/node/deleteNode",
         method: 'delete',
         data
     })
 }

// @Tags Node
// @Summary 删除Node
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Node"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /node/deleteNode [delete]
 export const deleteNodeByIds = (data) => {
     return service({
         url: "/node/deleteNodeByIds",
         method: 'delete',
         data
     })
 }

// @Tags Node
// @Summary 更新Node
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Node true "更新Node"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /node/updateNode [put]
 export const updateNode = (data) => {
     return service({
         url: "/node/updateNode",
         method: 'put',
         data
     })
 }


// @Tags Node
// @Summary 用id查询Node
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Node true "用id查询Node"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /node/findNode [get]
 export const findNode = (params) => {
     return service({
         url: "/node/findNode",
         method: 'get',
         params
     })
 }


// @Tags Node
// @Summary 分页获取Node列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Node列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /node/getNodeList [get]
 export const getNodeList = (params) => {
     return service({
         url: "/node/getNodeList",
         method: 'get',
         params
     })
 }