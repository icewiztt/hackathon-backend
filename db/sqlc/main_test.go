package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/thanhqt2002/hackathon/db/util"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("Cannot read config file", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}
	testQueries = New(conn)

	os.Exit(m.Run())
}
