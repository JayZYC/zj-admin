package util

import (
	"math/rand"
	"time"
	"unicode"
)

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
