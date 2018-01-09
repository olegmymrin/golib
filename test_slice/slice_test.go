package test_slice

import (
	"testing"
	"fmt"
)

func dummy(d []int) {
	d[0] = -1
}

func BenchmarkSample(_b *testing.B) {
	a := []int{1, 2, 3, 4}
	var b, c, d []int
	for i := 0; i < _b.N; i++ {
		b = a[:1]
		c = a[2:]
		d = append(b, c...)
		dummy(b)
		dummy(d)
	}
	fmt.Println(b, c, d)
}

func TestSubSlice(t *testing.T) {
	a := []int{1, 2, 3, 4}
	var b, c, d []int
	b = a[:1]
	fmt.Println(b)
	c = a[2:]
	fmt.Println(c)
	d = append(b, c...)
	fmt.Println(b, c, d)

}