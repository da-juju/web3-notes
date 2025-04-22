package result

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Resp 接口的统一返回值封装
type Resp struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"` // omitempty当data为空时,不展示这个字段
}

func OkWithoutData(ctx *gin.Context) {
	r := Resp{Code: http.StatusOK, Message: "操作成功", Data: nil}
	ctx.JSON(http.StatusOK, r)
}

func OkWithData(ctx *gin.Context, data interface{}) {
	r := Resp{Code: http.StatusOK, Message: "操作成功", Data: data}
	ctx.JSON(http.StatusOK, r)
}

func Error(ctx *gin.Context) {
	r := Resp{Code: http.StatusInternalServerError, Message: "操作失败", Data: nil}
	ctx.JSON(http.StatusInternalServerError, r)
}

func AjaxResult(ctx *gin.Context, rows int64) {
	if rows >= 1 {
		OkWithoutData(ctx)
	} else {
		Error(ctx)
	}
}
