package router

import (
	"github.com/gin-gonic/gin"
	v1 "zj-admin/router/v1"
	"zj-admin/util/middleware"
)

func Init(r *gin.Engine) {

	// 中间件注册
	r.Use(gin.Logger(), middleware.Auth(), middleware.Cors())

	api := r.Group("/api")

	v1.Register(api)

}
