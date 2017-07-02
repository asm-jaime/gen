package gen

import (
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
	// case 0
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
	// case negative
	{
		str := Str(-10)
		if str != "" {
			t.Error("Str should be empty")
		}
	}
	// case 0
	{
		str := Str(0)
		if str != "" {
			t.Error("Str should be empty")
		}
	}
	// case same strs
	{
		str1 := Str(10)
		str2 := Str(10)
		if str1 == "" {
			t.Error("Str empty")
		}
		if str1 == str2 {
			t.Error("Str generated the same string")
		}
	}
} // }}}

func TestStrNums(t *testing.T) { // {{{
	// case negative
	{
		str := StrNums(-10)
		if str != "" {
			t.Error("StrNums not empty for negative")
		}
	}
	// case 0
	{
		str := StrNums(0)
		if str != "" {
			t.Error("StrNums should be empty for 0")
		}
	}
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

func TestSliceRndGap(t *testing.T) {
	// case 2 elems
	{
		slice := SliceRndGap(2, 0, 10)
		// fmt.Println(slice)
		if len(slice) != 2 {
			t.Error("SliceRndGap  ")
		}
	}
	// case -10 elems
	{
		slice := SliceRndGap(-10, 0, 10)
		if len(slice) != 2 {
			t.Error("SliceRndGap ")
		}
	}
	// case 3 elems
	{
		slice := SliceRndGap(3, 0, 10)
		if len(slice) != 3 {
			t.Error("SliceRndGap ")
		}
	}
	// case 5 elems for 0-10
	{
		slice := SliceRndGap(6, 0, 10)
		if len(slice) != 2 {
			t.Error("SliceRndGap for overlap step not 2 elems")
		}
	}
	// case min==max
	{
		slice := SliceRndGap(0, 10, 10)
		if len(slice) != 2 {
			t.Error("SliceRndGap not 2 elems")
		}
	}
	// case with negative gap
	{
		slice := SliceRndGap(4, -12, 10)
		// fmt.Println(slice)
		if len(slice) != 4 {
			t.Error("SliceRndGap negative not 4 elems")
		}
	}
}
