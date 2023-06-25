package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
	"zj-admin/db"
	"zj-admin/util/response"
)

type LoginParams struct {
	Username string `json:"username" binding:"required"` // 用户名
	Password string `json:"password" binding:"required"` // 密码
}

var (
	ErrUserNotFound    = errors.New("用户不存在")
	ErrAccountDisabled = errors.New("该账户已停用")
)

func login(c *gin.Context) {
	var loginParams LoginParams
	err := c.ShouldBind(&loginParams)
	if err != nil {
		response.BadRequest(c, err)
		return
	}

	user, err := db.FindUserByNameAndPWD(loginParams.Username, loginParams.Password)
	if err != nil {
		response.BadRequest(c, ErrUserNotFound)
		return
	}

	if user.Status == -1 {
		response.BadRequest(c, ErrAccountDisabled)
		return
	}
}
