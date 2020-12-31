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

// @Tags DeptInfo
// @Summary 创建DeptInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DeptInfo true "创建DeptInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dept/createDeptInfo [post]
func CreateDeptInfo(c *gin.Context) {
	var dept model.DeptInfo
	_ = c.ShouldBindJSON(&dept)
	if err := service.CreateDeptInfo(dept); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags DeptInfo
// @Summary 删除DeptInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DeptInfo true "删除DeptInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dept/deleteDeptInfo [delete]
func DeleteDeptInfo(c *gin.Context) {
	var dept model.DeptInfo
	_ = c.ShouldBindJSON(&dept)
	if err := service.DeleteDeptInfo(dept); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags DeptInfo
// @Summary 批量删除DeptInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除DeptInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /dept/deleteDeptInfoByIds [delete]
func DeleteDeptInfoByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteDeptInfoByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags DeptInfo
// @Summary 更新DeptInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DeptInfo true "更新DeptInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /dept/updateDeptInfo [put]
func UpdateDeptInfo(c *gin.Context) {
	var dept model.DeptInfo
	_ = c.ShouldBindJSON(&dept)
	if err := service.UpdateDeptInfo(dept); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags DeptInfo
// @Summary 用id查询DeptInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DeptInfo true "用id查询DeptInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /dept/findDeptInfo [get]
func FindDeptInfo(c *gin.Context) {
	var dept model.DeptInfo
	_ = c.ShouldBindQuery(&dept)
	if err, redept := service.GetDeptInfo(dept.DeptID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"redept": redept}, c)
	}
}

// @Tags DeptInfo
// @Summary 分页获取DeptInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.DeptInfoSearch true "分页获取DeptInfo列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dept/getDeptInfoList [get]
func GetDeptInfoList(c *gin.Context) {
	var pageInfo request.DeptInfoSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetDeptInfoInfoList(pageInfo); err != nil {
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
