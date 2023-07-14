package jwt

import (
	"time"
	"zj-admin/model"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const TokenName = "Bearer"
const SigningKey = "zj-admin"

// 自定义Token过期时间
// 单位：秒
var TokenExp = time.Duration(600) * time.Second

// TokenExp token过期时间
// 单位：小时
var tokenExp = 24

// CustomClaims are custom claims extending default ones.
type CustomClaims struct {
	UserID         uuid.UUID `json:"user_id"`
	Username       string    `json:"username"`
	Nickname       string    `json:"nickname"`
	RoleID         uuid.UUID `json:"role_id"`
	OrganizationID uuid.UUID `json:"organization_id"` // 组织ID
	Phone          string    `json:"phone"`           //联系号码
	Email          string    `json:"email"`
	jwt.StandardClaims
}

func GetToken(user *model.User) (string, error) {

	// Set custom claims
	claims := &CustomClaims{
		UserID:         user.ID,
		Username:       user.UserName,
		Nickname:       user.NickName,
		RoleID:         user.RoleID,
		OrganizationID: user.OrganizationID,
		Phone:          user.Phone,
		Email:          user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(tokenExp) * time.Hour).Unix(),
		},
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SigningKey))
}

// GetUser 获取用户信息
func GetUser(ctx *gin.Context) CustomClaims {
	val, _ := ctx.Get("user")
	user := val.(*CustomClaims)
	return *user
}
