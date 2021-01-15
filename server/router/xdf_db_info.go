package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitDataBaseInfoRouter(Router *gin.RouterGroup) {
	DataBaseInfoRouter := Router.Group("DBInfo").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		DataBaseInfoRouter.POST("createDataBaseInfo", v1.CreateDataBaseInfo)   // 新建DataBaseInfo
		DataBaseInfoRouter.DELETE("deleteDataBaseInfo", v1.DeleteDataBaseInfo) // 删除DataBaseInfo
		DataBaseInfoRouter.DELETE("deleteDataBaseInfoByIds", v1.DeleteDataBaseInfoByIds) // 批量删除DataBaseInfo
		DataBaseInfoRouter.PUT("updateDataBaseInfo", v1.UpdateDataBaseInfo)    // 更新DataBaseInfo
		DataBaseInfoRouter.GET("findDataBaseInfo", v1.FindDataBaseInfo)        // 根据ID获取DataBaseInfo
		DataBaseInfoRouter.GET("getDataBaseInfoList", v1.GetDataBaseInfoList)  // 获取DataBaseInfo列表
	}
}
