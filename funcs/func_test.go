package funcs


import (
	"testing"
)

func call(f func(int) error) {
	f(0)
}

type Foo struct {
	t *testing.T
}

func (f *Foo) Bar(x int) error {
	f.t.Log("Bar")
	return nil
}

func TestMemberFunc(t *testing.T) {
	foo := Foo{t}
	call(foo.Bar)
}

func getNilFunc() func() {
	return nil
}

func TestNilFunc(t *testing.T) {
	getNilFunc()()
}