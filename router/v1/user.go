package v1

import (
	"zj-admin/db"
	"zj-admin/model"
	"zj-admin/util"
	"zj-admin/util/jwt"
	"zj-admin/util/pagination"
	"zj-admin/util/response"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func getUserInfo(c *gin.Context) {
	user := jwt.GetUser(c)
	perms := make([]string, 0)
	if !util.IsAdmin(user.RoleID) {
		var err error
		perms, err = db.FindPermByRoleID(user.RoleID)
		if err != nil {
			response.InternalServerError(c, err)
			return
		}
	}
	res := struct {
		UserName string    `json:"username"`
		Nickname string    `json:"nickname"`
		Mobile   string    `json:"mobile"`
		Email    string    `json:"email"`
		Perms    []string  `json:"perms"`
		Role     uuid.UUID `json:"role"`
	}{UserName: user.Username, Nickname: user.Nickname, Mobile: user.Phone, Email: user.Email, Perms: perms, Role: user.RoleID}

	response.OK(c, res)
}

func getUserList(c *gin.Context) {
	user := jwt.GetUser(c)

	var list pagination.PageResponse
	var err error

	page, err := pagination.BindPage(c)
	if err != nil {
		response.BadRequest(c, err)
		return
	}

	list, err = db.FindUserList(user, page)
	if err != nil {
		response.InternalServerError(c, err)
		return
	}

	response.OK(c, list)
}

func addUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBind(&user); err != nil {
		response.BadRequest(c, err)
		return
	}

	if db.UserIsExist(user.UserName) || db.PhoneExist(user.Phone) {
		response.BadRequest(c, response.ErrHasExist)
		return
	}

	if err := db.AddUser(user); err != nil {
		response.BadRequest(c, err)
		return
	}

	response.Success(c)

}

func updateUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBind(&user); err != nil {
		response.BadRequest(c, err)
		return
	}

	if err := db.UpdateUser(user); err != nil {
		response.InternalServerError(c, err)
		return
	}

	response.Success(c)
}

func deleteUser(c *gin.Context) {
	id := c.Query("id")
	uid, err := uuid.Parse(id)
	if err != nil {
		response.BadRequest(c, err)
		return
	}

	if err := db.DeleteUser(uid); err != nil {
		response.InternalServerError(c, err)
		return
	}

	response.Success(c)
}
