package test14

import (
	"fmt"
	"testing"
)

func TestSprintf(t *testing.T) {
	res := fmt.Sprintf("%d %s", 1, "%s%d%f")
	if res != "1 %s%d%f" {
		t.Fatal(res)
	}
}