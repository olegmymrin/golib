package test15

import (
	"testing"
)

type I interface {
	Foo()
}

func TestNil(t *testing.T) {
	var i I = nil
	if i != nil {
		t.Fail()
	}
}

type Impl int

func (i Impl) Foo() {
}

func TestImplNil(t *testing.T) {
	var i I = nil
	var v Impl
	i = v
	if v, ok := i.(Impl); !ok || v != 0 {
		t.Fail()
	}
}