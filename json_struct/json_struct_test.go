package json_struct

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"git.acronis.com/abm/policy-manager/kit"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
)

type JsonStruct struct {
	Val interface{}
}

func (s JsonStruct) Value() (driver.Value, error) {
	v, err := json.Marshal(s.Val)
	if err != nil {
		return nil, err
	}
	return driver.Value(v), nil
}

func (s *JsonStruct) Scan(src interface{}) error {
	if src == nil {
		return nil
	}
	fmt.Println("SRC", src, s)
	switch src.(type) {
	case string:
		return json.Unmarshal(([]byte)(src.(string)), s.Val)
	case []byte:
		return json.Unmarshal(src.([]byte), s.Val)
	default:
		return fmt.Errorf("Invalid row value type in database, Type: %v", reflect.TypeOf(src).String())
	}
}

type Embed struct {
	Name string
}

type MyStruct struct {
	Embed `sql:"json_struct;type:json" gorm:"embedded;embed_prefix:"`
	Id    string `json:"id"`
}

func (s MyStruct) TableName() string {
	return "NAMES"
}

func TestJsonStruct(t *testing.T) {
	s := MyStruct{
		Id:    `1`,
		Embed: Embed{`1`},
	}
	db, err := sql.Open("sqlite3", ":memory:")
	require.Nil(t, err)
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS NAMES (id  VARCHAR(255), json_struct VARCHAR(255))")
	require.Nil(t, err)
	orm, err := gorm.Open("sqlite", db)
	require.Nil(t, err)
	require.Nil(t, orm.Create(s).Error)
	got := MyStruct{}
	require.Nil(t, orm.First(&got).Error)
	require.Equal(t, s, got)
}

type S struct {
	A int         `json:"a"`
	O kit.RawJson `json:"o,omitempty"`
}

func TestOptional(t *testing.T) {
	v, err := json.Marshal(S{})
	require.Nil(t, err)
	require.Equal(t, `{"a":0}`, string(v))
}
