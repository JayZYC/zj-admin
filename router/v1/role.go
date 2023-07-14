package v1

import (
	"zj-admin/db"
	"zj-admin/model"
	"zj-admin/util/jwt"
	"zj-admin/util/pagination"
	"zj-admin/util/response"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func getRoleList(c *gin.Context) {
	user := jwt.GetUser(c)

	page, err := pagination.BindPage(c)
	if err != nil {
		response.BadRequest(c, err)
		return
	}

	result, err := db.FindRoleListByRoleID(user.RoleID, page)
	if err != nil {
		response.InternalServerError(c, err)
		return
	}

	response.OK(c, result)
}

func addRole(c *gin.Context) {
	var role model.Role
	if err := c.ShouldBind(&role); err != nil {
		response.BadRequest(c, err)
		return
	}

	user := jwt.GetUser(c)
	role.ParentID = user.RoleID

	err := db.AddRole(role)
	if err != nil {
		response.InternalServerError(c, err)
		return
	}

	response.Success(c)
}

func updateRole(c *gin.Context) {
	var role model.UpdateRole
	if err := c.ShouldBind(&role); err != nil {
		response.BadRequest(c, err)
		return
	}

	if err := db.UpdateRole(role); err != nil {
		response.InternalServerError(c, err)
		return
	}

	response.Success(c)
}

func deleteRole(c *gin.Context) {
	id := c.Query("id")
	uid, err := uuid.Parse(id)
	if err != nil {
		response.BadRequest(c, err)
		return
	}

	if err := db.DeleteRole(uid); err != nil {
		response.InternalServerError(c, err)
		return
	}

	response.Success(c)
}
