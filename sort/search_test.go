package sort

import (
	"sort"
	"testing"
)

func TestSearch1(t *testing.T) {
	values := []string{"a"}
	isA := func(i int) bool {
		t.Log(i)
		return values[i] == "a"
	}
	idx := sort.Search(len(values), isA)
	if idx == len(values) {
		t.Fatal("Not found")
	}
}

func TestSearch2(t *testing.T) {
	values := []string{"a", "b"}
	isA := func(i int) bool {
		t.Log(i)
		return values[i] == "a"
	}
	idx := sort.Search(len(values), isA)
	if idx == len(values) {
		t.Fatal("Not found")
	}
}

func TestSearch3(t *testing.T) {
	values := []string{"a", "b", "c"}
	isA := func(i int) bool {
		t.Log(i)
		return values[i] == "a"
	}
	idx := sort.Search(len(values), isA)
	if idx == len(values) {
		t.Fatal("Not found")
	}
}

func TestSearch4(t *testing.T) {
	values := []string{"a", "b", "c", "d"}
	isA := func(i int) bool {
		t.Log(i)
		return values[i] == "a"
	}
	idx := sort.Search(len(values), isA)
	if idx == len(values) {
		t.Fatal("Not found")
	}
}
