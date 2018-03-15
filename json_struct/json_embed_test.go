package json_struct

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

type VariableJson struct {
	json.RawMessage
	Type string `json:"type"`
}

func TestEmbedJsonStruct(t *testing.T) {
	var value VariableJson
	err := json.Unmarshal(([]byte)(`{"type":"1","value": 2}`), &value)
	require.Nil(t, err)
	require.Equal(t, "1", value.Type)

	data, err := json.Marshal(value)
	require.Nil(t, err)
	require.Equal(t, "1", value.Type)
}
