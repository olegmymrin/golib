package test10

import (
	"testing"
	"fmt"
	"runtime/debug"
)

func TestPanicExit(t *testing.T) {
	defer func(){
		fmt.Println("exit")
		err := recover()
		if err != nil {
			fmt.Println("panic:", err, string(debug.Stack()))
		}
	}()
	panic("panic1")
}
