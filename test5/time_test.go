package test5

import (
	"testing"
	"time"
	"reflect"
)

func TestTimeEquality(t *testing.T) {
	t1, err := time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2017-08-08 15:42:11 +0000 UTC")
	if err != nil {
		t.Fatal(err)
	}
	t2, err := time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2017-08-08 18:42:11 +0003 MSK")
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(t1, t2) {
		t.Fatalf("\n Expected: %v,\n      got: %v", t1, t2)
	}
}
