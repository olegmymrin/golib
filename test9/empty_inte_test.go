package test9

import "testing"

func f1(ei interface{}, t *testing.T) {
	f2(ei, t)
}

func f2(ei interface{}, t *testing.T) {
	if ei != nil {
		t.Fatal("Not nil!")
	}
}

func TestEmptyInterface(t *testing.T) {
	f1(nil, t)
}
