package response

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	ErrNoPermission    = errors.New("权限不足")
	ErrUserNotFound    = errors.New("账号或者密码错误")
	ErrAccountDisabled = errors.New("该账户已停用")
	ErrTokenExPire     = errors.New("登录过期")
	ErrOtherLogin      = errors.New("账号已在别处登录")
	ErrHasExist        = errors.New("账号或手机已存在")
	ErrParameter       = errors.New("参数错误")
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
		Message: "ok",
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
		Message: err.Error(),
	})
}

func NoPermission(c *gin.Context, err error) {
	c.JSON(http.StatusForbidden, CommonResponse{
		Message: err.Error(),
	})
}

func InternalServerError(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, CommonResponse{
		Message: err.Error(),
	})
}
