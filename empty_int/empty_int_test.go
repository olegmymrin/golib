package empty_int

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

type Optional struct {
	Field interface{} `json:"field"`
}

func TestEmptyInt(t *testing.T) {
	var o Optional
	require.Nil(t, json.Unmarshal(([]byte)("{}"), &o))
	require.Nil(t, o.Field)
	ser, err := json.Marshal(o)
	require.Nil(t, err)
	t.Log(t, string(ser))
	var o2 Optional
	require.Equal(t, o, o2)
}
