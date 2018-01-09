package test1

import (
	"testing"
	"fmt"
)

func Tes1tString(t *testing.T) {
	s := "0123"
	b := []byte(s)
	b[0] = "5"[0]
	s2 := string(b)
	fmt.Println(s, s2)
}
