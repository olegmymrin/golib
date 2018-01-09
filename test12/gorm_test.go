package golib

import (
	"database/sql"
	"testing"

	"github.com/jinzhu/gorm"

	_ "github.com/mattn/go-sqlite3"
)

type Row struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

func (r *Row) TableName() string {
	return "row"
}

func TestDistinct(t *testing.T) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatal(err)
	}

	if _, err := db.Exec("create table row (id int, name varchar(255))"); err != nil {
		t.Fatal(err)
	}

	query, err := gorm.Open("sqlite", db)
	if err != nil {
		t.Fatal(err)
	}

	if err := query.Create(&Row{1, "1"}).Error; err != nil {
		t.Fatal(err)
	}
	if err := query.Create(&Row{1, "1"}).Error; err != nil {
		t.Fatal(err)
	}
	if err := query.Create(&Row{2, "2"}).Error; err != nil {
		t.Fatal(err)
	}

	var res []Row
	query.LogMode(true)
	rows := query.Model(&Row{}).Where("id = 1").Select("distinct row.*").Find(&res)
	if rows.Error != nil {
		t.Fatal(rows.Error)
	}
	query.LogMode(false)
	t.Log(res)
}
