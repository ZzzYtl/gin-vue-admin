package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitDBAPI(Router *gin.RouterGroup) {
	ApiRouter := Router.Group("db_api").Use(middleware.OperationRecord())
	{
		ApiRouter.GET("/cluster/select", v1.ClusterSelect)
	}
}
