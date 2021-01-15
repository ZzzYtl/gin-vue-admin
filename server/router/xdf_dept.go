package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitDeptInfoRouter(Router *gin.RouterGroup) {
	DeptInfoRouter := Router.Group("dept").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		DeptInfoRouter.POST("createDeptInfo", v1.CreateDeptInfo)   // 新建DeptInfo
		DeptInfoRouter.DELETE("deleteDeptInfo", v1.DeleteDeptInfo) // 删除DeptInfo
		DeptInfoRouter.DELETE("deleteDeptInfoByIds", v1.DeleteDeptInfoByIds) // 批量删除DeptInfo
		DeptInfoRouter.PUT("updateDeptInfo", v1.UpdateDeptInfo)    // 更新DeptInfo
		DeptInfoRouter.GET("findDeptInfo", v1.FindDeptInfo)        // 根据ID获取DeptInfo
		DeptInfoRouter.GET("getDeptInfoList", v1.GetDeptInfoList)  // 获取DeptInfo列表
		DeptInfoRouter.POST("getAllDepts", v1.GetAllDepts)
	}
}
