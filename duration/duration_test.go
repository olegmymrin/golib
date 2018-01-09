package duration

import (
	"encoding/json"
	"testing"
	"time"
)

type dur struct {
	Dur time.Duration `json:"dur"`
}

func TestDurationJson(t *testing.T) {
	v := dur{}
	err := json.Unmarshal([]byte(`{"dur" : "3m1s"}`), &v)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%v", v.Dur)
}