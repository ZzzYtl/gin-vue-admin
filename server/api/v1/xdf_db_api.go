package v1

import (
	"gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ClusterSelect(c *gin.Context) {
	result := response.InitResult()
	DbName:= c.Query("db_name")

	data, err := service.QueryClusterInfo(DbName)
	if err != nil {
		result.SetParams(1, err.Error(), nil)
	} else {
		result.SetParams(0, "", data)
	}

	c.JSON(http.StatusOK, result)
	return
}
