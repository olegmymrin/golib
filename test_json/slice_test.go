package test_json

import (
	"testing"
	"encoding/json"
)

func TestVarEmptySliceMarshal(t *testing.T) {
	var s1 []string
	res, _ := json.Marshal(s1)
	if (string(res) != "null") {
		t.Fatal(string(res))
	}
}

func TestInitEmptySliceMarshal(t *testing.T) {
	s1 := []string{}
	res, _ := json.Marshal(s1)
	if (string(res) != "[]") {
		t.Fatal(string(res))
	}
}

func TestMakeEmptySliceMarshal(t *testing.T) {
	s1 := make([]string, 0)
	res, _ := json.Marshal(s1)
	if (string(res) != "[]") {
		t.Fatal(string(res))
	}
}