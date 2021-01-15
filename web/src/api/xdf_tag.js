import service from '@/utils/request'

// @Tags Node
// @Summary 获取Tags
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Node true "获取Tags"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /node/createNode [post]
export const getAllTags = (data) => {
     return service({
         url: "/tag/getAllTags",
         method: 'post',
         data
     })
 }
