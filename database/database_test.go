package database_test

import (
	"os"
	"testing"

	"github.com/ochom/sdp-lib/database"
)

func initDB(t *testing.T) database.Repo {

	os.Setenv("DATABASE_DNS", "db.sqlite")
	i, err := database.New(database.SQLite)
	if err != nil {
		t.Fatal(err)
	}

	return i
}
