package test_regex

import (
	"regexp"
	"testing"
	"github.com/stretchr/testify/require"
)

func TestRegexp(t *testing.T) {
	re, err := regexp.Compile("[0-9]")
	require.Nil(t, err)
	v := re.FindAll([]byte("456"))

}
