package maps

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRangeNil(t *testing.T) {
	var m map[string]bool
	for i := range m {
		t.Log(m[i])
	}
}

func TestAppendToSlice(t *testing.T) {
	v1 := "1"
	v2 := "2"
	v3 := "3"
	values := []string{v1, v2, v3}
	result := map[string][]*string{}
	for i := range values {
		row := *(&values[i])
		result["values"] = append(result["values"], &row)
	}
	expected := map[string][]*string{
		"values": []*string{&v1, &v2, &v3},
	}
	require.Equal(t, expected, result)
}
