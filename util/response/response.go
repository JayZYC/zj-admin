package response

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CommonResponse struct {
	Success bool        `json:"success"` // 调用结果是否成功
	Message string      `json:"msg"`     // 消息;
	Data    interface{} `json:"data"`    // 实际数据
}

func New(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, CommonResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func Success(c *gin.Context) {
	c.JSON(http.StatusOK, CommonResponse{
		Success: true,
		Message: "OK",
	})
}

func OK(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, CommonResponse{
		Success: true,
		Message: "ok",
		Data:    data,
	})
}

func Message(c *gin.Context, message string) {
	c.JSON(http.StatusOK, CommonResponse{
		Success: true,
		Message: message,
	})
}

func BadRequest(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, CommonResponse{
		Message: err.Error(),
	})
}

func Fail(c *gin.Context, err error) {
	c.JSON(http.StatusOK, CommonResponse{
		Message: err.Error(),
	})
}

func NotFound(c *gin.Context, err error) {
	c.JSON(http.StatusNotFound, CommonResponse{
		Message: fmt.Sprintf("%v", err),
	})
}

func NoPermission(c *gin.Context, err error) {
	c.JSON(http.StatusForbidden, CommonResponse{
		Message: fmt.Sprintf("%v", err),
	})
}

func InternalServerError(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, CommonResponse{
		Message: fmt.Sprintf("%v", err),
	})
}
