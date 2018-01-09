package orms

import (
	"testing"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/require"
	_ "github.com/mattn/go-sqlite3"
)

type Base struct {
	ID string
	Name string
}

type Child struct {
	Base
	Value string
}

func TestNestedStructs(t *testing.T) {
	db, err := gorm.Open("sqlite", "sqlite3", ":memory:")
	require.Equal(t, nil, err)
	require.Equal(t, nil, db.Exec("create table children (id int primary key, name varhchar(255), value varhchar(255))").Error)
	v := Child{
		Base: Base{
			ID: "1",
			Name: "1",
		},
		Value: "1",
	}
	require.Equal(t, nil, db.Model(&Child{}).Create(&v).Error)
	var got Child
	require.Equal(t, nil, db.Model(&Child{}).First(&got, &v).Error)
	require.Equal(t, v, got)
}