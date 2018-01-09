package test11

import (
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	repeats := 5
	timeout := 1
	started := time.Now()
	elapsed := time.Since(started)
	for repeats - (int)(elapsed / (time.Duration(timeout) * time.Second)) > 0 {
		time.Sleep(950 * time.Millisecond)
		elapsed = time.Since(started)
	}
}