package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitSentinelClusterRouter(Router *gin.RouterGroup) {
	NodeRouter := Router.Group("sentinel_cluster").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		NodeRouter.POST("getAllSentinelClusters", v1.GetAllSentinelClusters)
	}
}
