package dbproducts

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/djudju12/order-system/ms-products/configs"
	_ "github.com/lib/pq"
)

var (
	testQueries *Queries
	testDB      *sql.DB
)

func TestMain(m *testing.M) {
	config, err := configs.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot read configuration:", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot open db connection:", err)
	}

	testQueries = New(testDB)
	os.Exit(m.Run())
}
