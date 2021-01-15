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

// @Tags DataBaseInfo
// @Summary 创建DataBaseInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DataBaseInfo true "创建DataBaseInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /DBInfo/createDataBaseInfo [post]
func CreateDataBaseInfo(c *gin.Context) {
	var DBInfo model.DataBaseInfo
	_ = c.ShouldBindJSON(&DBInfo)
	if err := service.CreateDataBaseInfo(DBInfo); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags DataBaseInfo
// @Summary 删除DataBaseInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DataBaseInfo true "删除DataBaseInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /DBInfo/deleteDataBaseInfo [delete]
func DeleteDataBaseInfo(c *gin.Context) {
	var DBInfo model.DataBaseInfo
	_ = c.ShouldBindJSON(&DBInfo)
	if err := service.DeleteDataBaseInfo(DBInfo); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags DataBaseInfo
// @Summary 批量删除DataBaseInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除DataBaseInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /DBInfo/deleteDataBaseInfoByIds [delete]
func DeleteDataBaseInfoByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteDataBaseInfoByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags DataBaseInfo
// @Summary 更新DataBaseInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DataBaseInfo true "更新DataBaseInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /DBInfo/updateDataBaseInfo [put]
func UpdateDataBaseInfo(c *gin.Context) {
	var DBInfo model.DataBaseInfo
	_ = c.ShouldBindJSON(&DBInfo)
	if err := service.UpdateDataBaseInfo(DBInfo); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags DataBaseInfo
// @Summary 用id查询DataBaseInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DataBaseInfo true "用id查询DataBaseInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /DBInfo/findDataBaseInfo [get]
func FindDataBaseInfo(c *gin.Context) {
	var DBInfo model.DataBaseInfo
	_ = c.ShouldBindQuery(&DBInfo)
	if err, reDBInfo := service.GetDataBaseInfo(DBInfo.ClusterID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reDBInfo": reDBInfo}, c)
	}
}

// @Tags DataBaseInfo
// @Summary 分页获取DataBaseInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.DataBaseInfoSearch true "分页获取DataBaseInfo列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /DBInfo/getDataBaseInfoList [get]
func GetDataBaseInfoList(c *gin.Context) {
	var pageInfo request.DataBaseInfoSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetDataBaseInfoInfoList(pageInfo); err != nil {
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
