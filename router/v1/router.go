package v1

import "github.com/gin-gonic/gin"

func Register(api *gin.RouterGroup) {

	const versionPrefix = "/v1"

	a := api.Group(versionPrefix)

	/** auth 权限验证 **/

	a.POST("/login", login)

	//a.POST("/login/captcha", loginByCaptcha)
	//
	//a.POST("/login/wechat", loginByWechat)
	//
	//a.POST("/login/ak", loginByAK)
	//
	//a.POST("/logout", logout)

}
