package gen

import (
	"fmt"
	"strings"
	"testing"
)

func TestStr(t *testing.T) { // {{{
	str1 := Str(10)
	str2 := Str(10)
	//fmt.Println(str1)
	//fmt.Println(str2)
	if str1 == "" {
		t.Error("Str empty")
	}
	if str1 == str2 {
		t.Error("Str generated the same string")
	}
} // }}}

func TestStrNums(t *testing.T) { // {{{
	str1 := StrNums(10)
	str2 := StrNums(10)
	//fmt.Println(str1)
	//fmt.Println(str2)
	if str1 == "" {
		t.Error("StrNums empty")
	}
	if str1 == str2 {
		t.Error("StrNums generated the same string")
	}

	str1_sp := strings.Split(str1, ",")
	if str1_sp[0] == "0" {
		t.Error("StrNums string started with 0")
	}
} // }}}

func TestMinMax(t *testing.T) { // {{{
	// case 1
	{
		st := MinMax(0, 10)
		fmt.Println(st)
		if (st > 10) || (st < 0) {
			t.Error("MinMax not within")
		}
	}
	// case 2
	{
		st := MinMax(-1, 1)
		fmt.Println(st)
		if (st > 1) || (st < -1) {
			t.Error("MinMax not within")
		}
	}
	// case 3
	{
		st := MinMax(1, -1)
		fmt.Println(st)
		if (st > 1) || (st < -1) {
			t.Error("MinMax not within")
		}
	}
} // }}}

func TestRandToken(t *testing.T) { // {{{
	token1 := Token(10)
	token2 := Token(10)
	fmt.Println(token1)
	fmt.Println(token2)
	if token1 == "" {
		t.Error("Token empty")
	}
	if token1 == token2 {
		t.Error("Token generated the same string")
	}
} // }}}
