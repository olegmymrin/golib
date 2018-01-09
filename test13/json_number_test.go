package test13

import (
	"testing"
	"encoding/json"
	"bytes"
)

type TaskConfig struct {
	Priority          string
	HeartBeatInterval uint32
	Timeout           int32
	MaxAttempts       uint32
}

type Config struct {
	TaskSettings TaskConfig
}

func TestUseNumber(t *testing.T) {
	buf := bytes.NewBufferString(`
	{
		"TaskSettings": {
			"Priority":          "normal",
			"HeartBeatInterval": 60,
			"Timeout":           60,
			"MaxAttempts":       60
		}
	}`)
	d := json.NewDecoder(buf)
	//d.UseNumber()

	c := Config{}
	if err := d.Decode(&c); err != nil {
		t.Fatal(err)
	}

	if c.TaskSettings.MaxAttempts != 60 {
		t.Fatal(c)
	}
}
