package test_requrie

import(
	"testing"
	"github.com/stretchr/testify/require"
)

func TestSlices(t *testing.T) {
	require.Equal(t, []string{"1","2"}, []string{"2","1"})
}
