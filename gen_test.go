package gen

import (
	// "fmt"
	"strings"
	"testing"
)

func TestIAbs(t *testing.T) { // {{{
	// case negative
	{
		num := IAbs(-10)
		if num != 10 {
			t.Error("IAbs not correct for -10")
		}
	}
	// case zero
	{
		num := IAbs(0)
		if num != 0 {
			t.Error("IAbs not correct for 0")
		}
	}
	// case positive
	{
		num := IAbs(10)
		if num != 10 {
			t.Error("IAbs not correct for 10")
		}
	}
} // }}}

func TestStr(t *testing.T) { // {{{
	// case same
	{
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
	}
	// case zero
	{
		str := Str(0)
		if str != "" {
			t.Error("Str should be empty")
		}
	}
	// case negative
	{
		// str := Str(-10)
		// fmt.Println(str)
	}
} // }}}

func TestStrNums(t *testing.T) { // {{{
	// case same
	{
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
	}
	// case first digit should not be 0
	{
		str := StrNums(10)
		str_sp := strings.Split(str, ",")
		if str_sp[0] == "0" {
			t.Error("StrNums gened fdigit as 0")
		}
	}
	// case negative
	{
		// str := StrNums(-10)
		// fmt.Println(str)
	}
} // }}}

func TestMinMax(t *testing.T) { // {{{
	// case 1
	{
		st := MinMax(0, 10)
		// fmt.Println(st)
		if (st > 10) || (st < 0) {
			t.Error("MinMax not within")
		}
	}
	// case 2
	{
		st := MinMax(-1, 1)
		// fmt.Println(st)
		if (st > 1) || (st < -1) {
			t.Error("MinMax not within")
		}
	}
	// case 3
	{
		st := MinMax(1, -1)
		// fmt.Println(st)
		if (st > 1) || (st < -1) {
			t.Error("MinMax not within")
		}
	}
} // }}}

func TestRandTokenB64(t *testing.T) { // {{{
	// case 1
	{
		token := TokenB64(10)
		// fmt.Println(token)
		if token == "" {
			t.Error("Token empty")
		}
	}
	// case negative
	{
		// token := TokenB64(-10)
	}
} // }}}
