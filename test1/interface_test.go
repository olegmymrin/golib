package test1

import(
	"testing"
)

//type I interface {
//	foo()
//}
//
//type B struct {
//}
//
//func (b *B) foo() {
//	fmt.Println("My name is B")
//}
//
//func getB() interface{} {
//	return B{}
//}
//
//
//func Tes1tInterface(t *testing.T) {
//	//ei := getB()
//	//if b, ok := ei.(B); ok {
//	//	var i I = &b
//	//	i.foo()
//	//}
//	//wg := sync.WaitGroup{}
//	//wg.Add(3)
//	c := make(chan bool, 1)
//	for i := 0; i < 3; i++ {
//		i := i
//		go func() {
//			fmt.Println(i)
//			time.Sleep(1)
//			fmt.Println("<<", i)
//			c<- true
//			//wg.Done()
//		}()
//	}
//	for i := 0; i < 3; i++ {
//		select {
//		case <-c:
//		}
//	}
//	//wg.Done()
//	fmt.Println("Done")
//}
//
//func Tes1tNilInterface(t *testing.T) {
//	var i I
//	i.foo()
//}
//
//func foo(x interface{}) bool {
//	_, ok := x.([]interface{})
//	return ok
//}
//
//
//func Tes1tCast(t *testing.T) {
//	t.Log(foo([]int{1}))
//}
//
//func Tes1tSlice(t *testing.T) {
//	s1 := []int{1}
//	s2 := append(s1, 2)
//	s2[0] = 3
//	if s1[0] != 3 {
//		t.Fatal("diff")
//	}
//}

type A struct{}
func (*A) foo(int) {}

type B struct{ A }
func (*B) foo(string) {}

type I interface {
	foo(int)
}

func TestInher(t *testing.T) {
	b := B{}
	b.foo(1)
	var i I = &B{}
	i.foo(1)
}