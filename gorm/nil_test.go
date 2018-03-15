package gorm_test

import (
	"database/sql/driver"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/require"
	_ "github.com/ziutek/mymysql/godrv"
)

type Valuer int64

func (v Valuer) Value() (driver.Value, error) {
	return int64(v), nil
}

func (v *Valuer) Scan(src interface{}) error {
	if src == nil {
		return nil
	}
	*v = Valuer(src.(int64))
	return nil
}

type ObjectWithNil struct {
	ID    string
	Value *Valuer
}

func TestCreateNilSqlite(t *testing.T) {
	db, err := gorm.Open("sqlite3", ":memory:")
	require.Nil(t, err)
	require.Nil(t, db.CreateTable(&ObjectWithNil{}).Error)
	err = db.Create(&ObjectWithNil{}).Error
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestCreateNilMysql(t *testing.T) {
	db, err := gorm.Open("mysql", "root:qwe123QWE@tcp(localhost:3306)/policy_manager_test?parseTime=true&multiStatements=true")
	require.Nil(t, err)
	err = db.Create(&ObjectWithNil{ID: "1"}).Error
	if err != nil {
		t.Fatal(err.Error())
	}
}
