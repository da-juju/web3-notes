package controller

import (
	"github.com/gin-gonic/gin"
	"goweb/controller/result"
)

// Upload 单文件上次
func Upload(c *gin.Context) {
	file, _ := c.FormFile("file")
	err := c.SaveUploadedFile(file, "./file/"+file.Filename)
	if err != nil {
		result.Error(c)
	}
	result.OkWithData(c, "文件上次成功")
}

// MutiUpload 多文件上传
func MutiUpload(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["file"]
	for _, file := range files {
		_ = c.SaveUploadedFile(file, "./file/"+file.Filename)
	}
	result.OkWithData(c, "文件上次成功")
}
