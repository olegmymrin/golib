package method_set

import (
	"testing"
)

type I interface {
	Get() int
	Set(i int)
}

type V struct {
	v int
}

func (v V) Get() int {
	return v.v
}

func (v *V) Set(i int) {
	v.v = i
}

func TestMethodSet(t *testing.T) {
	var i I = &V{}
	t.Log(i.Get())
}
