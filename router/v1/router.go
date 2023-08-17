package v1

import "github.com/gin-gonic/gin"

func Register(api *gin.RouterGroup) {

	const versionPrefix = "/v1"

	a := api.Group(versionPrefix)

	/** auth 权限验证 **/

	a.POST("/login", login)

	//a.POST("/logout", logout)

	/** user 用户 **/

	a.GET("/user/info", getUserInfo)

	a.GET("/users", getUserList)

	a.POST("/user", addUser)

	a.PUT("/user", updateUser)

	a.DELETE("/user", deleteUser)

	/** role 角色 **/

	a.GET("/roles", getRoleList)

	a.POST("/role", addRole)

	a.PUT("/role", updateRole)

	a.DELETE("/role", deleteRole)

	/** organization 组织 **/

	a.GET("/organizations", getOrgList)

	a.POST("/organization", addOrg)

	a.PUT("/organization", updateOrg)

	a.DELETE("/organization", deleteOrg)

	/** menu 菜单 **/

	a.GET("/menu", getMenu)

	a.POST("/menu", addMenu)

	a.PUT("/menu", updateMenu)

	a.DELETE("/menu", deleteMenu)

	a.GET("/menu/routes", getMenuTree)

}
