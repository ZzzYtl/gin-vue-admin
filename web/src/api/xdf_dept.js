import service from '@/utils/request'

// @Tags DeptInfo
// @Summary 创建DeptInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DeptInfo true "创建DeptInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dept/createDeptInfo [post]
export const createDeptInfo = (data) => {
     return service({
         url: "/dept/createDeptInfo",
         method: 'post',
         data
     })
 }


// @Tags DeptInfo
// @Summary 删除DeptInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DeptInfo true "删除DeptInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dept/deleteDeptInfo [delete]
 export const deleteDeptInfo = (data) => {
     return service({
         url: "/dept/deleteDeptInfo",
         method: 'delete',
         data
     })
 }

// @Tags DeptInfo
// @Summary 删除DeptInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除DeptInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dept/deleteDeptInfo [delete]
 export const deleteDeptInfoByIds = (data) => {
     return service({
         url: "/dept/deleteDeptInfoByIds",
         method: 'delete',
         data
     })
 }

// @Tags DeptInfo
// @Summary 更新DeptInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DeptInfo true "更新DeptInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /dept/updateDeptInfo [put]
 export const updateDeptInfo = (data) => {
     return service({
         url: "/dept/updateDeptInfo",
         method: 'put',
         data
     })
 }


// @Tags DeptInfo
// @Summary 用id查询DeptInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DeptInfo true "用id查询DeptInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /dept/findDeptInfo [get]
 export const findDeptInfo = (params) => {
     return service({
         url: "/dept/findDeptInfo",
         method: 'get',
         params
     })
 }


// @Tags DeptInfo
// @Summary 分页获取DeptInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取DeptInfo列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dept/getDeptInfoList [get]
 export const getDeptInfoList = (params) => {
     return service({
         url: "/dept/getDeptInfoList",
         method: 'get',
         params
     })
 }