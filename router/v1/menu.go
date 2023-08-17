package v1

import (
	"strconv"
	"zj-admin/db"
	"zj-admin/model"
	"zj-admin/util"
	"zj-admin/util/jwt"
	"zj-admin/util/response"

	"github.com/gin-gonic/gin"
)

func getMenu(c *gin.Context) {
	user := jwt.GetUser(c)
	menus, err := db.FindMenu(user.RoleID)
	if err != nil {
		response.InternalServerError(c, err)
		return
	}
	response.OK(c, menus)
}

func getMenuTree(c *gin.Context) {
	user := jwt.GetUser(c)
	menus, err := db.FindMenu(user.RoleID)
	if err != nil {
		response.InternalServerError(c, err)
		return
	}

	menu := util.GetMenuTree(0, menus)

	response.OK(c, menu)
}

func addMenu(c *gin.Context) {
	var Menu model.Menu
	if err := c.ShouldBind(&Menu); err != nil {
		response.BadRequest(c, err)
		return
	}

	err := db.AddMenu(Menu)
	if err != nil {
		response.InternalServerError(c, err)
		return
	}

	response.Success(c)
}

func updateMenu(c *gin.Context) {
	var Menu model.UpdateMenu
	if err := c.ShouldBind(&Menu); err != nil {
		response.BadRequest(c, err)
		return
	}

	if err := db.UpdateMenu(Menu); err != nil {
		response.InternalServerError(c, err)
		return
	}

	response.Success(c)
}

func deleteMenu(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.BadRequest(c, err)
		return
	}

	if err := db.DeleteMenu(id); err != nil {
		response.InternalServerError(c, err)
		return
	}

	response.Success(c)
}
