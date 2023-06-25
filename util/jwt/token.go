package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"time"
	"zj-admin/model"
)

const TokenName = "Bearer"
const SigningKey = "zj-admin"

// TokenExp token过期时间
// 单位：小时
var TokenExp = 24

// CustomClaims are custom claims extending default ones.
type CustomClaims struct {
	UserID         uuid.UUID `json:"user_id"`
	Username       string    `json:"username"`
	RoleID         uuid.UUID `json:"role_id"`
	OrganizationID uuid.UUID `json:"organization_id" bson:"organization_id"` // 组织ID
	Phone          string    `json:"phone" bson:"phone"`                     //联系号码
	jwt.StandardClaims
}

func GetToken(user *model.User) (string, error) {

	// Set custom claims
	claims := &CustomClaims{
		UserID:         user.ID,
		Username:       user.UserName,
		RoleID:         user.RoleID,
		OrganizationID: user.OrganizationID,
		Phone:          user.Phone,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(TokenExp) * time.Hour).Unix(),
		},
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SigningKey))
}
