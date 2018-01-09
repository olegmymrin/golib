package db_nil

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
)

func TestNilParam(t *testing.T) {
	db, err := sql.Open("sqlite3", ":memory:")
	require.Nil(t, err)
	_, err = db.Exec("CREATE TABLE ids (id INT, name VARHCHAR(255))")
	require.Nil(t, err)
	_, err = db.Exec("INSERT INTO ids VALUES (1, NULL)")
	require.Nil(t, err)
	res, err := db.Query("SELECT id FROM IDS where name is ?", nil)
	require.Nil(t, err)
	defer res.Close()
	require.True(t, res.Next())
	var id int
	err = res.Scan(&id)
	require.Nil(t, err)
	require.Equal(t, 1, id)
}
