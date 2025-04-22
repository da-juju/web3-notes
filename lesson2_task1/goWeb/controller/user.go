package controller

import (
	"github.com/gin-gonic/gin"
	"goweb/controller/result"
	"goweb/models"
	"goweb/service"
	"log"
	"net/http"
)

func Get(c *gin.Context) {
	param := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"id": param,
	})
}

func Insert(c *gin.Context) {
	var user models.User
	// json: Unmarshal(non-pointer models.User)==>json.Unmarshal 函数要求目标变量必须是一个指针
	// err := c.ShouldBindJSON(user)
	err := c.ShouldBindJSON(&user)
	if err != nil {
		log.Fatalf("user save err: %v", err)
		return
	}
	row := service.InsertUser(&user)
	result.AjaxResult(c, row)
}
