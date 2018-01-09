package maps

import (
	"testing"
)

func TestRangeNil(t *testing.T) {
	var m map[string]bool
	for i := range m {
		t.Log(m[i])
	}
}
