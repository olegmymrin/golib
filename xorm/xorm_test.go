package main

import (
	"testing"

	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
)

func TestQueryInterface(t *testing.T) {
	xorm.NewEngineGroup()
	engine, err := xorm.NewEngine("sqlite3", ":memory:")
	require.Nil(t, err)
	engine.QueryInterface()
}
