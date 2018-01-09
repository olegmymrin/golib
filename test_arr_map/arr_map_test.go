package test_arr_map

import (
	"testing"
)

func TestArrayMap(t *testing.T) {
	m := map[[8]byte]string{
		[8]byte{0,1,2,3,4,5,6,7}: "1",
	}
	t.Log(m)
}
