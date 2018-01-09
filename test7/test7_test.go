package test7

import (
	"testing"
	"path"
)

func TestPath(t *testing.T) {
	join := path.Join("a", "b")
	if join != "a/b" {
		t.Fatal(join)
	}
}