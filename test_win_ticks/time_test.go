package test_win_ticks

import (
	"testing"
	"time"
	"fmt"
)

func TestNanoseconds(t *testing.T) {
	now := time.Now().UTC()
	fmt.Println("now        ", now, "s", now.Unix(), "ns", now.UnixNano())
	now = now.AddDate(245, 0, 0)
	fmt.Println("now + 245", now, "s", now.Unix(), "ns", now.UnixNano())
	t.Fail()
}

func timediff() time.Duration {
    t0 := time.Now()
    for {
        t := time.Now()
        if t != t0 {
            return t.Sub(t0)
        }
    }
}

func TestGetTickCount(t *testing.T) {
	fmt.Println(timediff())
}