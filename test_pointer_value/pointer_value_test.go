package test_pointer_value

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func TestPointerValue(t *testing.T) {
	var p = new(int)
	*p = 0
	v := *p
	v = 1
	require.Equal(t, 0, *p)
	require.Equal(t, 1, v)
}