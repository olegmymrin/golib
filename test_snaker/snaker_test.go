package test_snaker

import (
	"testing"
	"github.com/iancoleman/strcase"
	"github.com/stretchr/testify/require"
)

func TestSnaker(t *testing.T) {
	require.Equal(t, "hostId", strcase.ToLowerCamel("host_id"))
}