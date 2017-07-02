package gen

import (
	"encoding/base64"
	"math/rand"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var numberRunes = []rune("0123456789")

// IAbs absolute value of n as a int number
func IAbs(n int) int { // {{{
	if n < 0 {
		return -n
	}
	if n == 0 {
		return 0
	}
	return n
} // }}}

// Str string of length n as a array of runes
func Str(n uint) string { // {{{
	rand.Seed(time.Now().UTC().UnixNano())
	str := make([]rune, n)
	for i := range str {
		str[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(str)
} // }}}

// StrNums string of length n as a number
func StrNums(n int) string { // {{{
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
} // }}}

// MinMax num within a range
func MinMax(min, max int) int { // {{{
	rand.Seed(time.Now().UTC().UnixNano())
	if min > max {
		buf := min
		min = max
		max = buf
	}
	return rand.Intn(max-min) + min
} // }}}

// TokenB64 random string of length n as a base64 token
func TokenB64(n uint) string { // {{{
	rand.Seed(time.Now().UTC().UnixNano())
	some_bytes := make([]byte, n)
	rand.Read(some_bytes)
	return base64.StdEncoding.EncodeToString(some_bytes)
} // }}}

// SliceRndGap return slice with random points in the interim gap
func SliceRndGap(len int, min int, max int) (res []int) {
	if len < 1 {
		return res
	}
	if min > max {
		buf := min
		min = max
		max = buf
	}
	step := IAbs((max - min) / len)
	if step < 1 {
		return res
	}
	res = append(res, min)
	start := min + 1
	end := start + step - 1
	for i := 0; i < len; i++ {
		res = append(res, MinMax(start, end))
		start = end + 2
		end = start + step - 1
	}
	return res
}
