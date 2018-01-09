package virt

import (
	"fmt"
	"testing"
)

type A interface {
	Foo()
}

type str interface {
	getStr() string
}

type a struct {
	str
}

func (this *a) Foo() {
	fmt.Println(this.getStr())
}

type s struct{}

func (this *s) getStr() string {
	return "default str"
}

type b struct {
	*a
}

func (this *b) getStr() string {
	return "b str"
}

type valueStr struct{}

func (this valueStr) getStr() string {
	return "valueStr"
}

func TestVirtual(t *testing.T) {
	//	a := &a{&s{}}
	//	b := &b{a}
	//	a.str = b
	//	b.Foo()
	var str str
	str = valueStr{}
	fmt.Println(str.getStr())
}

func addToSlice(sl []string) {
	sl[0] = "1"
}

func TestSlice(t *testing.T) {
	sl := []string{"0"}
	addToSlice(sl)
	fmt.Println(sl)
}
