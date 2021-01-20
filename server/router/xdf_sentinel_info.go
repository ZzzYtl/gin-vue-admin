package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitSentinelInfoRouter(Router *gin.RouterGroup) {
	SentinelInfoRouter := Router.Group("sentinelinfo").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		SentinelInfoRouter.POST("createSentinelInfo", v1.CreateSentinelInfo)   // 新建SentinelInfo
		SentinelInfoRouter.DELETE("deleteSentinelInfo", v1.DeleteSentinelInfo) // 删除SentinelInfo
		SentinelInfoRouter.DELETE("deleteSentinelInfoByIds", v1.DeleteSentinelInfoByIds) // 批量删除SentinelInfo
		SentinelInfoRouter.PUT("updateSentinelInfo", v1.UpdateSentinelInfo)    // 更新SentinelInfo
		SentinelInfoRouter.GET("findSentinelInfo", v1.FindSentinelInfo)        // 根据ID获取SentinelInfo
		SentinelInfoRouter.GET("getSentinelInfoList", v1.GetSentinelInfoList)  // 获取SentinelInfo列表
	}
}
