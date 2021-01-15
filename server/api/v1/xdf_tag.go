package v1

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Tags Node
// @Summary 获取Tag列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.NodeSearch true "获取Tag列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /node/getNodeList [get]
func GetAllTags(c *gin.Context) {
	if err, list, total := service.GetTagList(); err != nil {
	    global.GVA_LOG.Error("获取失败", zap.Any("err", err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
        }, "获取成功", c)
    }
}
