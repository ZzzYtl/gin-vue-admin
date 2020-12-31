package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitBackUpDBRouter(Router *gin.RouterGroup) {
	BackUpDBRouter := Router.Group("BackDB").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		BackUpDBRouter.POST("createBackUpDB", v1.CreateBackUpDB)   // 新建BackUpDB
		BackUpDBRouter.DELETE("deleteBackUpDB", v1.DeleteBackUpDB) // 删除BackUpDB
		BackUpDBRouter.DELETE("deleteBackUpDBByIds", v1.DeleteBackUpDBByIds) // 批量删除BackUpDB
		BackUpDBRouter.PUT("updateBackUpDB", v1.UpdateBackUpDB)    // 更新BackUpDB
		BackUpDBRouter.GET("findBackUpDB", v1.FindBackUpDB)        // 根据ID获取BackUpDB
		BackUpDBRouter.GET("getBackUpDBList", v1.GetBackUpDBList)  // 获取BackUpDB列表
	}
}
