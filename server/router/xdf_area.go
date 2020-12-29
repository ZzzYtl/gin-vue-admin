package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitAreaRouter(Router *gin.RouterGroup) {
	AreaRouter := Router.Group("area").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		AreaRouter.POST("createArea", v1.CreateArea)   // 新建Area
		AreaRouter.DELETE("deleteArea", v1.DeleteArea) // 删除Area
		AreaRouter.DELETE("deleteAreaByIds", v1.DeleteAreaByIds) // 批量删除Area
		AreaRouter.PUT("updateArea", v1.UpdateArea)    // 更新Area
		AreaRouter.GET("findArea", v1.FindArea)        // 根据ID获取Area
		AreaRouter.GET("getAreaList", v1.GetAreaList)  // 获取Area列表
	}
}
