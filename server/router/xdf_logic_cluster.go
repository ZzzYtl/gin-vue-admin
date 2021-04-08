package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitLogicClusterRouter(Router *gin.RouterGroup) {
	NodeRouter := Router.Group("logicCluster").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		NodeRouter.POST("getAllLogicClusters", v1.GetAllLogicClusters)
	}
}

