package v1

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Tags SentinelDBClusterInfo
// @Summary 创建SentinelDBClusterInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SentinelDBClusterInfo true "创建SentinelDBClusterInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /SentinelDBCluster/createSentinelDBClusterInfo [post]
func CreateSentinelDBClusterInfo(c *gin.Context) {
	var SentinelDBCluster model.SentinelDBClusterInfo
	_ = c.ShouldBindJSON(&SentinelDBCluster)
	if err := service.CreateSentinelDBClusterInfo(SentinelDBCluster); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

func ChangeOldMaster2Slave(c *gin.Context) {
	var SentinelDBCluster service.OldMasterForm
	_ = c.ShouldBindJSON(&SentinelDBCluster)
	if err := service.ChangeOldMaster2Slave(SentinelDBCluster); err != nil {
		global.GVA_LOG.Error("变更失败!", zap.Any("err", err))
		response.FailWithMessage("变更失败", c)
	} else {
		response.OkWithMessage("变更成功", c)
	}
}

// @Tags SentinelDBClusterInfo
// @Summary 删除SentinelDBClusterInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SentinelDBClusterInfo true "删除SentinelDBClusterInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /SentinelDBCluster/deleteSentinelDBClusterInfo [delete]
func DeleteSentinelDBClusterInfo(c *gin.Context) {
	var SentinelDBCluster model.SentinelDBClusterInfo
	_ = c.ShouldBindJSON(&SentinelDBCluster)
	if err := service.DeleteSentinelDBClusterInfo(SentinelDBCluster); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags SentinelDBClusterInfo
// @Summary 批量删除SentinelDBClusterInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SentinelDBClusterInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /SentinelDBCluster/deleteSentinelDBClusterInfoByIds [delete]
func DeleteSentinelDBClusterInfoByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteSentinelDBClusterInfoByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags SentinelDBClusterInfo
// @Summary 更新SentinelDBClusterInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SentinelDBClusterInfo true "更新SentinelDBClusterInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /SentinelDBCluster/updateSentinelDBClusterInfo [put]
func UpdateSentinelDBClusterInfo(c *gin.Context) {
	var SentinelDBCluster model.SentinelDBClusterInfo
	_ = c.ShouldBindJSON(&SentinelDBCluster)
	if err := service.UpdateSentinelDBClusterInfo(SentinelDBCluster); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags SentinelDBClusterInfo
// @Summary 用id查询SentinelDBClusterInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SentinelDBClusterInfo true "用id查询SentinelDBClusterInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /SentinelDBCluster/findSentinelDBClusterInfo [get]
func FindSentinelDBClusterInfo(c *gin.Context) {
	var SentinelDBCluster model.SentinelDBClusterInfo
	_ = c.ShouldBindQuery(&SentinelDBCluster)
	if err, reSentinelDBCluster := service.GetSentinelDBClusterInfo(SentinelDBCluster.ClusterID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reSentinelDBCluster": reSentinelDBCluster}, c)
	}
}

// @Tags SentinelDBClusterInfo
// @Summary 分页获取SentinelDBClusterInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SentinelDBClusterInfoSearch true "分页获取SentinelDBClusterInfo列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /SentinelDBCluster/getSentinelDBClusterInfoList [get]
func GetSentinelDBClusterInfoList(c *gin.Context) {
	var pageInfo request.SentinelDBClusterInfoSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetSentinelDBClusterInfoInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
