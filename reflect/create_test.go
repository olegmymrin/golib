package reflect_test

import (
	"reflect"
	"testing"
)

type Object struct {
}

func TestCreate(t *testing.T) {
	typ := reflect.TypeOf(&Object{}).Elem()
	value := reflect.New(typ).Interface().(*Object)
	t.Log(value)
}
