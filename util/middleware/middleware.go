package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"zj-admin/cache"
	jwtToken "zj-admin/util/jwt"
	"zj-admin/util/response"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v9"
	"github.com/rs/zerolog/log"
)

var (
	skipper = []string{"login"} // 跳过鉴权的路由
)

// Auth  token校验
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {

		for _, s := range skipper {
			ok := strings.Contains(c.Request.URL.Path, s)
			if ok {
				c.Next()
				return
			}
		}

		tokenString := strings.Split(c.Request.Header.Get("Authorization"), " ")
		if len(tokenString) != 2 || tokenString[0] != "Bearer" {
			response.BadRequest(c, response.ErrTokenExPire)
			c.Abort()
			return
		}

		token, err := jwt.ParseWithClaims(tokenString[1], &jwtToken.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtToken.SigningKey), nil
		})
		if err != nil {
			response.BadRequest(c, response.ErrTokenExPire)
			c.Abort()
			return
		}

		claims, ok := token.Claims.(*jwtToken.CustomClaims)
		if !ok {
			response.BadRequest(c, fmt.Errorf("please relogin"))
			c.Abort()
			return
		}

		val, err := cache.Get(fmt.Sprintf("%s%s", cache.Token, claims.UserID.String()))
		if err == redis.Nil {
			response.BadRequest(c, response.ErrTokenExPire)
			c.Abort()
			return
		}

		if err != nil {
			response.InternalServerError(c, err)
			c.Abort()
			return
		}

		if val != tokenString[1] {
			response.InternalServerError(c, response.ErrOtherLogin)
			c.Abort()
			return
		}

		cache.Expire(fmt.Sprintf("%s%s", cache.Token, claims.UserID.String()), jwtToken.TokenExp)

		c.Set("user", claims)
		c.Next()
	}

}

// Cors CORS（跨域资源共享）中间件
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin") //请求头部
		if origin != "" {
			//接收客户端发送的origin （重要！）
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			//服务器支持的所有跨域请求的方法
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
			//允许跨域设置可以返回其他子段，可以自定义字段
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			// 允许浏览器（客户端）可以解析的头部 （重要）
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
			//设置缓存时间
			c.Header("Access-Control-Max-Age", "172800")
			//允许客户端传递校验信息比如 cookie (重要)
			c.Header("Access-Control-Allow-Credentials", "true")
		}

		//允许类型校验
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "ok!")
		}

		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic info is: %v", err)
			}
		}()
		c.Next()
	}
}
