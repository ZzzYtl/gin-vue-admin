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

type CreateSentinelInfoForm struct {
	ClusterName string `json:"cluster_name"`
	Sentinels []model.SentinelInfo `json:"sentinels"`
}

// @Tags SentinelInfo
// @Summary 创建SentinelInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SentinelInfo true "创建SentinelInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sentinelinfo/createSentinelInfo [post]
func CreateSentinelInfo(c *gin.Context) {
	//var sentinelinfo model.SentinelInfo
	//_ = c.ShouldBindJSON(&sentinelinfo)
	//if err := service.CreateSentinelInfo(sentinelinfo); err != nil {
    //    global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
	//	response.FailWithMessage("创建失败", c)
	//} else {
	//	response.OkWithMessage("创建成功", c)
	//}

	var form CreateSentinelInfoForm
	_ = c.ShouldBindJSON(&form)

	sentinelClusterInfo := model.SentinelClusterInfo{
		Name:  form.ClusterName,
	}
	sentinelClusterId, err := service.CreateSentinelClusterInfo(sentinelClusterInfo);
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	}

	for _, v := range form.Sentinels {
		v.SentinelClusterID = sentinelClusterId
		if err := service.CreateSentinelInfo(v); err != nil {
			global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
			response.FailWithMessage("创建失败", c)
			break
		}
	}
	response.OkWithMessage("创建成功", c)

}

// @Tags SentinelInfo
// @Summary 删除SentinelInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SentinelInfo true "删除SentinelInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sentinelinfo/deleteSentinelInfo [delete]
func DeleteSentinelInfo(c *gin.Context) {
	var sentinelinfo model.SentinelInfo
	_ = c.ShouldBindJSON(&sentinelinfo)
	findErr := global.GVA_DB.Where("sentinel_id = ?", sentinelinfo.SentinelID).First(&sentinelinfo).Error
	if err := service.DeleteSentinelInfo(sentinelinfo); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		if findErr == nil {
			db := global.GVA_DB.Model(&model.SentinelDBClusterInfo{})
			var total int64
			db = db.Where("`sentinel_cluster_id` = ?", sentinelinfo.SentinelClusterID)
			countErr := db.Count(&total).Error
			if countErr == nil {
				if total == 0 {
					global.GVA_DB.Delete(&model.SentinelClusterInfo{SentinelClusterID: sentinelinfo.SentinelClusterID})
				}
			}
		}
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags SentinelInfo
// @Summary 批量删除SentinelInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SentinelInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /sentinelinfo/deleteSentinelInfoByIds [delete]
func DeleteSentinelInfoByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteSentinelInfoByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags SentinelInfo
// @Summary 更新SentinelInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SentinelInfo true "更新SentinelInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sentinelinfo/updateSentinelInfo [put]
func UpdateSentinelInfo(c *gin.Context) {
	var sentinelinfo model.SentinelInfo
	_ = c.ShouldBindJSON(&sentinelinfo)
	if err := service.UpdateSentinelInfo(sentinelinfo); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags SentinelInfo
// @Summary 用id查询SentinelInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SentinelInfo true "用id查询SentinelInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sentinelinfo/findSentinelInfo [get]
func FindSentinelInfo(c *gin.Context) {
	var sentinelinfo model.SentinelInfo
	_ = c.ShouldBindQuery(&sentinelinfo)
	if err, resentinelinfo := service.GetSentinelInfo(sentinelinfo.SentinelID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"resentinelinfo": resentinelinfo}, c)
	}
}

// @Tags SentinelInfo
// @Summary 分页获取SentinelInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SentinelInfoSearch true "分页获取SentinelInfo列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sentinelinfo/getSentinelInfoList [get]
func GetSentinelInfoList(c *gin.Context) {
	var pageInfo request.SentinelInfoSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetSentinelInfoInfoList(pageInfo); err != nil {
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
