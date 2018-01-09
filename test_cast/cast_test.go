package test_cast

import (
	"testing"
)

type base struct {
}

type parent struct {
	base
	name string
}

func TestNetstedCast(t *testing.T) {
	var p interface{} = parent{}
	if _, ok := p.(base); !ok {
		t.Fatal("Cannot cast to nested")
	}
}

func TestNetstedPointerCast(t *testing.T) {
	var p interface{} = &parent{}
	if _, ok := p.(*base); !ok {
		t.Fatal("Cannot cast to nested")
	}
}

//func TestCastToEmptyIntSlice(t *testing.T) {
//	res := []string{"1"}
//	([]interface{})(res)
//}

func CallInt(fn func(interface{})) {

}

func TestInterfaceCastInFunc(t *testing.T) {
	CallInt(func (i int) {
		return
	})
}