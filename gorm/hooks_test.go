package gorm_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
)

type Embed struct {
	ValueS string `gorm:"-"`
}

type Value struct {
	Embed
	Value int64
}

func (v *Value) BeforeSave() (err error) {
	fmt.Println("BEFORE_SAVE")
	v.Value, err = strconv.ParseInt(v.Embed.ValueS, 10, 64)
	return
}

func (v *Value) BeforeCreate() error {
	fmt.Println("BEFORE_CREATE")
	return v.BeforeSave()
}

func (v *Value) AfterFind() error {
	fmt.Println("AFTER_FIND")
	v.Embed.ValueS = strconv.FormatInt(v.Value, 10)
	return nil
}

func TestBeforeSave(t *testing.T) {
	db, err := gorm.Open("sqlite3", ":memory:")
	require.Nil(t, err)
	require.Nil(t, db.Exec(`create table "values" (value bigint not null primary key)`).Error)
	require.Nil(t, db.Save(&Value{Embed: Embed{"42"}}).Error)
	got := Value{}
	require.Nil(t, db.First(&got).Error)
	require.Equal(t, int64(42), got.Value, "%#v", got)
}
