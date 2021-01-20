package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitSentinelDBClusterInfoRouter(Router *gin.RouterGroup) {
	SentinelDBClusterInfoRouter := Router.Group("SentinelDBCluster").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		SentinelDBClusterInfoRouter.POST("createSentinelDBClusterInfo", v1.CreateSentinelDBClusterInfo)   // 新建SentinelDBClusterInfo
		SentinelDBClusterInfoRouter.DELETE("deleteSentinelDBClusterInfo", v1.DeleteSentinelDBClusterInfo) // 删除SentinelDBClusterInfo
		SentinelDBClusterInfoRouter.DELETE("deleteSentinelDBClusterInfoByIds", v1.DeleteSentinelDBClusterInfoByIds) // 批量删除SentinelDBClusterInfo
		SentinelDBClusterInfoRouter.PUT("updateSentinelDBClusterInfo", v1.UpdateSentinelDBClusterInfo)    // 更新SentinelDBClusterInfo
		SentinelDBClusterInfoRouter.GET("findSentinelDBClusterInfo", v1.FindSentinelDBClusterInfo)        // 根据ID获取SentinelDBClusterInfo
		SentinelDBClusterInfoRouter.GET("getSentinelDBClusterInfoList", v1.GetSentinelDBClusterInfoList)  // 获取SentinelDBClusterInfo列表
	}
}
