package json_struct_opt

import (
	"encoding/json"
	"testing"

	"git.acronis.com/abm/policy-manager/kit"
	"github.com/stretchr/testify/require"
)

type S struct {
	A int          `json:"a"`
	O *kit.RawJson `json:"o,omitempty"`
}

func TestOptional(t *testing.T) {
	v, err := json.Marshal(S{})
	require.Nil(t, err)
	require.Equal(t, `{"a":0}`, string(v))
}
