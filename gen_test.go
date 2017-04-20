package gen

import (
	"fmt"
	"testing"
)

func TestStr(t *testing.T) { // {{{
	str1 := Str(10)
	str2 := Str(10)
	fmt.Println(str1)
	fmt.Println(str2)
	if str1 == "" {
		t.Error("Str empty")
	}
	if str1 == str2 {
		t.Error("Str generated same string")
	}
} // }}}
