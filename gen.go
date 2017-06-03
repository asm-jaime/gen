package gen

import (
	"math/rand"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var numberRunes = []rune("0123456789")

// Str retrun a string with n length letters
func Str(n int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	str := make([]rune, n)
	for i := range str {
		str[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(str)
}

// Str retrun a string with n length numbers
func StrNums(n int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	strn := make([]rune, n)
	for i := range strn {
		strn[i] = numberRunes[rand.Intn(len(numberRunes))]
	}

	if strn[0] == '0' {
		cutRunes := []rune("123456789")
		strn[0] = cutRunes[rand.Intn(len(cutRunes))]
	}
	return string(strn)
}
