package v1

import (
	"zj-admin/db"
	"zj-admin/model"
	"zj-admin/util/jwt"
	"zj-admin/util/response"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func getOrgList(c *gin.Context) {
	user := jwt.GetUser(c)

	result, err := db.FindOrgListByOrgID(user.OrganizationID)
	if err != nil {
		response.InternalServerError(c, err)
		return
	}

	response.OK(c, result)
}

func addOrg(c *gin.Context) {
	var org model.Organization
	if err := c.ShouldBind(&org); err != nil {
		response.BadRequest(c, err)
		return
	}

	err := db.AddOrg(org)
	if err != nil {
		response.InternalServerError(c, err)
		return
	}

	response.Success(c)
}

func updateOrg(c *gin.Context) {
	var org model.UpdateOrg
	if err := c.ShouldBind(&org); err != nil {
		response.BadRequest(c, err)
		return
	}

	if err := db.UpdateOrg(org); err != nil {
		response.InternalServerError(c, err)
		return
	}

	response.Success(c)
}

func deleteOrg(c *gin.Context) {
	id := c.Query("id")
	uid, err := uuid.Parse(id)
	if err != nil {
		response.BadRequest(c, err)
		return
	}

	if err := db.DeleteOrg(uid); err != nil {
		response.InternalServerError(c, err)
		return
	}

	response.Success(c)
}
