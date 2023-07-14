package util

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"regexp"
	"strings"
	"time"
	"unicode"

	"github.com/google/uuid"
)

const admin = "5138e5da-4161-485a-bd1f-5f5a10de2f80"

// SpecialLetters
// 如果存在特殊字符，直接在特殊字符前添加\
// 判断是否为字母： unicode.IsLetter(v)
// 判断是否为十进制数字： unicode.IsDigit(v)
// 判断是否为数字： unicode.IsNumber(v)
// 判断是否为空白符号： unicode.IsSpace(v)
// 判断是否为Unicode标点字符 :unicode.IsPunct(v)
// 判断是否为中文：unicode.Han(v)
func SpecialLetters(letter rune) (bool, []rune) {
	if unicode.IsPunct(letter) || unicode.IsSymbol(letter) || unicode.Is(unicode.Han, letter) {
		var chars []rune
		chars = append(chars, '\\', letter)
		return true, chars
	}
	return false, nil
}

// Shuffle 数组洗牌
func Shuffle[T any](arr []T) []T {
	slc := make([]T, len(arr))
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for i, randIndex := range r.Perm(len(arr)) {
		slc[i] = arr[randIndex]
	}
	return slc
}

func IsMD5(str string) bool {
	str = strings.ToLower(str)
	md5Reg := regexp.MustCompile(`^[a-f0-9]{32}$`)
	return md5Reg.MatchString(str)
}

func Md5encode(arg0 string) string {
	h := md5.New()
	h.Write([]byte(arg0))
	cipherStr := h.Sum(nil)
	return strings.ToLower(hex.EncodeToString(cipherStr)) // 输出加密结果
}

// IsAdmin 判断是否为超管
func IsAdmin(id uuid.UUID) bool {
	return id.String() == admin
}

// func Where(c *gin.Context, db *gorm.DB) *gorm.DB {
// 	user := jwt.GetUser(c)
// 	if IsAdmin(user.RoleID) {
// 		return db
// 	}
// 	db = db.Where()
// }
