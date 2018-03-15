package inherit_const

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

type Enum string

const (
	Enum1 Enum = "1"
	Enum2 Enum = "2"
)

func TestConstTypeInherit(t *testing.T) {
	require.Equal(t, reflect.TypeOf(Enum1), reflect.TypeOf(Enum2))
}
