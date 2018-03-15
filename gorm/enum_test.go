package gorm_test

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/require"
)

type Enum string

const (
	Enum1 Enum = "1"
	Enum2 Enum = "2"
)

type Object struct {
	ID    string
	Value Enum
}

func TestGormEnumsSqlite(t *testing.T) {
	db, err := gorm.Open("sqlite3", ":memory:")
	require.Nil(t, err)
	require.Nil(t, db.CreateTable(&Object{}).Error)
	require.Nil(t, db.Create(&Object{
		ID:    "1",
		Value: Enum1,
	}).Error)
}

func TestGormEnumsMysql(t *testing.T) {
	db, err := gorm.Open("mysql", "root:qwe123QWE@tcp(localhost:3306)/policy_manager_test?parseTime=true&multiStatements=true")
	require.Nil(t, err)
	require.Nil(t, db.Create(&Object{
		ID:    "1",
		Value: Enum1,
	}).Error)
}
