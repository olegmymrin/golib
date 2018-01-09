package map_equal

import (
	"testing"
	"reflect"
)

func TestMapEqual(t *testing.T) {
	x := make(map[string][]string)
	//x["a"] = []string{"b"}
	var y map[string][]string
	//y["a"] = []string{"b"}
	if (!reflect.DeepEqual(x, y)) {
		t.Fatal("Different", x, y)
	}
}
