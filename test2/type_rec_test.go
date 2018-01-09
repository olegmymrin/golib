package test2

import  (
	"fmt"
	"testing"
)

//type I1 interface {
//	I2
//	f1()
//}
//
//type I2 interface {
//	I3
//	f2()
//}
//
//type I3 interface {
//	I1
//	f3()
//}

type A struct {
	b *B
}

type B struct {
	//a *A
}

func TestTypeRecursion(t *testing.T) {
	a := &A{}
	b := &B{}
	a.b = b
	fmt.Printf("%v", a)
}