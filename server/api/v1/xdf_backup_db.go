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

// @Tags BackUpDB
// @Summary 创建BackUpDB
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BackUpDB true "创建BackUpDB"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /BackDB/createBackUpDB [post]
func CreateBackUpDB(c *gin.Context) {
	var BackDB model.BackUpDB
	_ = c.ShouldBindJSON(&BackDB)
	if err := service.CreateBackUpDB(BackDB); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags BackUpDB
// @Summary 删除BackUpDB
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BackUpDB true "删除BackUpDB"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /BackDB/deleteBackUpDB [delete]
func DeleteBackUpDB(c *gin.Context) {
	var BackDB model.BackUpDB
	_ = c.ShouldBindJSON(&BackDB)
	if err := service.DeleteBackUpDB(BackDB); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags BackUpDB
// @Summary 批量删除BackUpDB
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除BackUpDB"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /BackDB/deleteBackUpDBByIds [delete]
func DeleteBackUpDBByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteBackUpDBByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags BackUpDB
// @Summary 更新BackUpDB
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BackUpDB true "更新BackUpDB"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /BackDB/updateBackUpDB [put]
func UpdateBackUpDB(c *gin.Context) {
	var BackDB model.BackUpDB
	_ = c.ShouldBindJSON(&BackDB)
	if err := service.UpdateBackUpDB(BackDB); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags BackUpDB
// @Summary 用id查询BackUpDB
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BackUpDB true "用id查询BackUpDB"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /BackDB/findBackUpDB [get]
func FindBackUpDB(c *gin.Context) {
	var BackDB model.BackUpDB
	_ = c.ShouldBindQuery(&BackDB)
	if err, reBackDB := service.GetBackUpDB(BackDB.BackUpID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reBackDB": reBackDB}, c)
	}
}

// @Tags BackUpDB
// @Summary 分页获取BackUpDB列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.BackUpDBSearch true "分页获取BackUpDB列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /BackDB/getBackUpDBList [get]
func GetBackUpDBList(c *gin.Context) {
	var pageInfo request.BackUpDBSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetBackUpDBInfoList(pageInfo); err != nil {
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
