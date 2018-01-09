package int_casts

import (
	"testing"
	"fmt"
)

type Reader interface {
	Read()
}

type reader1 struct {
}

func (*reader1) Read() {
	fmt.Println("reader1")
}

type reader2 struct {
}

func (*reader2) Read() {
	fmt.Println("reader2")
}

func ReadMe(r *Reader) {
	*r = &reader2{}
}

func TestIntCast(t *testing.T) {
	var r1 = &reader1{}
	ReadMe(&r1)
	r1.Read()
}

