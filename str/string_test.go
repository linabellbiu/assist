package str

import (
	"fmt"
	"testing"
)

func TestJoin(t *testing.T) {
	var p_uint uint = 1
	var p_uint8 uint8 = 2
	var p_uint16 uint16 = 111
	var p_uint32 uint32 = 323232
	var p_uint64 uint64 = 646464646
	var p_int int = 646464646
	var p_int64 int64 = 646464646
	var p_string string = "dfdfd"
	fmt.Println(Join(p_uint, p_uint8, p_uint16, p_uint32, p_uint64, p_int, p_int64, p_string))
}

func TestJoinError(t *testing.T) {
	var p_uint uint = 1
	var p_uint8 uint8 = 2
	var p_uint16 uint16 = 111
	var p_uint32 uint32 = 323232
	var p_uint64 uint64 = 646464646
	var p_int int = 646464646
	var p_int64 int64 = 646464646
	var p_string string = "dfdfd"
	var p_struct = struct{}{}
	fmt.Println(Join(p_uint, p_uint8, p_uint16, p_uint32, p_uint64, p_int, p_int64, p_string, p_struct))
}
