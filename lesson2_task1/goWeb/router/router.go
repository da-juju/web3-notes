package router

import (
	"github.com/gin-gonic/gin"
	"goweb/controller"
	"net/http"
)

// SetupRouter 设置路由
func SetupRouter() *gin.Engine {
	ginServer := gin.Default()

	// 人员的路由分组
	userGroup := ginServer.Group("/user")
	{
		userGroup.GET("/get/:id", controller.Get)
		userGroup.POST("/insert", controller.Insert)
	}

	fileGroup := ginServer.Group("/file")
	{
		fileGroup.POST("/upload", controller.Upload)
		fileGroup.POST("mutipload", controller.MutiUpload)
	}

	// 部门的路由分组
	deptGroup := ginServer.Group("/dept")
	{
		deptGroup.GET("/get", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"id": 123, "deptName": "研发组"})
		})
		deptGroup.POST("/add", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "新增部门成功"})
		})
	}

	return ginServer
}
