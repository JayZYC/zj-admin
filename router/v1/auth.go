package v1

import (
	"fmt"
	"zj-admin/cache"
	"zj-admin/db"
	"zj-admin/util"
	"zj-admin/util/jwt"
	"zj-admin/util/response"

	"github.com/gin-gonic/gin"
)

type LoginParams struct {
	Username string `json:"username" form:"username" binding:"required"` // 用户名
	Password string `json:"password" form:"password" binding:"required"` // 密码
}

func login(c *gin.Context) {
	var loginParams LoginParams
	err := c.ShouldBind(&loginParams)
	if err != nil {
		response.BadRequest(c, err)
		return
	}

	if !util.IsMD5(loginParams.Password) {
		response.BadRequest(c, response.ErrUserNotFound)
		return
	}

	user, err := db.FindUserByNameAndPWD(loginParams.Username, loginParams.Password)
	if err != nil {
		response.BadRequest(c, response.ErrUserNotFound)
		return
	}

	if !util.IsAdmin(user.RoleID) {
		// 查询并缓存用户按钮权限
		perm, err := db.FindPermByRoleID(user.RoleID)
		if err != nil {
			response.InternalServerError(c, err)
			return
		}
		if err := cache.SetArr(fmt.Sprintf("%s%s", cache.Perm, user.ID.String()), perm, 0); err != nil {
			response.InternalServerError(c, err)
			return
		}
	}

	if user.Status == -1 {
		response.BadRequest(c, response.ErrAccountDisabled)
		return
	}

	token, err := jwt.GetToken(user)
	if err != nil {
		response.InternalServerError(c, err)
		return
	}

	// 缓存自定义过期token
	if err := cache.Set(fmt.Sprintf("%s%s", cache.Token, user.ID.String()), token, jwt.TokenExp); err != nil {
		response.InternalServerError(c, err)
		return
	}

	response.OK(c, map[string]string{"accessToken": token, "tokenType": jwt.TokenName})
}
