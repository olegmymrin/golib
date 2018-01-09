package test4

import (
	"encoding/json"
	"testing"
)

type B struct {
	json.RawMessage
}

type Data struct {
	A int `json:"a"`
	B B `json:"b"`
}

func TestJsonRawMessage(t *testing.T) {
	value := `{"a": 1, "b": { "b1": "b1" }}`
	var d Data
	if err := json.Unmarshal(([]byte)(value), &d); err != nil {
		t.Fatal(err)
	}
	msg, err := json.Marshal(d)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(msg))
}