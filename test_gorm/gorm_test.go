package test_gorm

import (
	"database/sql"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/require"
	_ "github.com/mattn/go-sqlite3"
	"database/sql/driver"
	"fmt"
	"reflect"
)

type Reader interface {
	Read()
}

type readerImpl struct {
	Name string
}

func (r *readerImpl) Read() {

}

func TestLoadToInteface(t *testing.T) {
	db, err := sql.Open("sqlite3", ":memory:")
	require.Equal(t, nil, err)
	orm, err := gorm.Open("sqlite", db)
	require.Equal(t, nil, err)
	err = orm.CreateTable(readerImpl{}).Error
	require.Equal(t, nil, err)
	err = orm.Create(readerImpl{
		Name: "1",
	}).Error
	var reader Reader
	err = orm.Model(readerImpl{}).Find(&reader).Error
	require.Equal(t, nil, err)
}

func TestDeleteByInteface(t *testing.T) {
	db, err := sql.Open("sqlite3", ":memory:")
	require.Equal(t, nil, err)
	orm, err := gorm.Open("sqlite3", db)
	require.Equal(t, nil, err)
	err = orm.CreateTable(readerImpl{}).Error
	require.Equal(t, nil, err)
	err = orm.Create(readerImpl{
		Name: "1",
	}).Error
	var reader Reader = &readerImpl{
		Name: "1",
	}
	err = orm.Model(readerImpl{}).Delete(reader).Error
	require.Equal(t, nil, err)
	var res []readerImpl
	err = orm.Model(readerImpl{}).Find(&res).Error
	require.Equal(t, nil, err)
	require.Equal(t, []readerImpl{}, res)
}

func TestUpdateByInteface(t *testing.T) {
	db, err := sql.Open("sqlite3", ":memory:")
	require.Equal(t, nil, err)
	orm, err := gorm.Open("sqlite3", db)
	require.Equal(t, nil, err)
	err = orm.CreateTable(readerImpl{}).Error
	require.Equal(t, nil, err)
	err = orm.Create(readerImpl{
		Name: "1",
	}).Error
	var reader Reader = &readerImpl{
		Name: "2",
	}
	err = orm.Model(readerImpl{}).Model(readerImpl{}).Update(reader).Error
	require.Equal(t, nil, err)
	var res []readerImpl
	err = orm.Model(readerImpl{}).Find(&res).Error
	require.Equal(t, nil, err)
	v := reader.(*readerImpl)
	require.Equal(t, []readerImpl{*v}, res)
}

func TestFindSlicePointers(t *testing.T) {
	db, err := sql.Open("sqlite3", ":memory:")
	require.Equal(t, nil, err)
	orm, err := gorm.Open("sqlite3", db)
	require.Equal(t, nil, err)
	err = orm.CreateTable(readerImpl{}).Error
	require.Equal(t, nil, err)
	v := readerImpl{
		Name: "1",
	}
	err = orm.Create(&v).Error
	var res []*readerImpl
	err = orm.Model(readerImpl{}).Find(&res).Error
	require.Equal(t, nil, err)
	require.Equal(t, []*readerImpl{&v}, res)
}

type Parent struct {
	string `sql:"column:name"`
}

func (p Parent) Value() (driver.Value, error) {
	return driver.Value(p.string), nil
}

func (p *Parent) Scan(src interface{}) error {
	if src == nil {
		return nil
	}
	switch src.(type) {
	case string:
		p.string = src.(string)
	case []byte:
		p.string = (string)(src.([]byte))
	default:
		return fmt.Errorf("Invalid row value type in database, Type: %v", reflect.TypeOf(src).String())
	}
	return nil
}

type Child struct {
	Parent `gorm:"embedded;embedded_prefix:"`
}

type value struct {
	Name Child `gorm:"embedded;embedded_prefix:child_"`
}

func TestCreateNestedStruct(t *testing.T) {
	db, err := sql.Open("sqlite3", ":memory:")
	require.Equal(t, nil, err)
	orm, err := gorm.Open("sqlite3", db)
	require.Equal(t, nil, err)
	err = orm.CreateTable(value{}).Error
	require.Equal(t, nil, err)
	v := value{Name: Child{Parent{"1"}},}
	err = orm.Create(&v).Error
	require.Equal(t, nil, err)
	var res []value
	err = orm.Model(value{}).Find(&res).Error
	require.Equal(t, nil, err)
	require.Equal(t, []value{v}, res)
}

