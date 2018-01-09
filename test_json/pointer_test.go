package test_json


import (
	"testing"
	"encoding/json"
	"github.com/stretchr/testify/require"
)

type PP struct{
	I *int `json:"int,omitempty"`
}

func TestOmitPointer(t *testing.T) {
	i := 0
	v, err := json.Marshal(&PP{&i})
	require.Nil(t, err)
	require.Equal(t, []byte(`{"int":0}`), v)
}