package gorm_test

import (
	"os"
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
)

func TestGormRelations(t *testing.T) {
	type Field struct {
		ID       string
		ObjectID string
		Name     string
	}
	type Object struct {
		ID     string
		Fields []Field `gorm:"ForeignKey:ObjectID"`
	}
	os.Remove("test.db3")
	db, err := gorm.Open("sqlite3", "test.db3")
	require.Nil(t, err)
	db.CreateTable(&Object{}, &Field{})
	o := Object{
		ID: "1",
		Fields: []Field{
			{
				ID:   "1",
				Name: "Field1",
			},
			{
				ID:   "2",
				Name: "Field2",
			},
		},
	}
	require.Nil(t, db.Create(&o).Error)
	got := Object{
		ID: "1",
	}
	require.Nil(t, db.Model(&got).Related(&got.Fields).Select(&got, &got).Error)
	require.Equal(t, o, got)
	toDelete := Object{
		ID: "1",
	}
	require.Nil(t, db.Model(&toDelete).Related(&got.Fields).Delete(&toDelete, &toDelete).Error)
	objects := []Object{}
	require.Nil(t, db.Find(&objects).Error)
	require.Len(t, objects, 0)
	fields := []Field{}
	require.Nil(t, db.Find(&fields).Error)
	require.Len(t, fields, 0)
}
