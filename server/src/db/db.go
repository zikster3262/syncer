package db

import (
	"concurrency/src/utils"
	"os"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	SQLXConnection *sqlx.DB
	err            error
)

type SQLxProvider struct{}

var onceSQLx sync.Once

func OpenSQLx() *sqlx.DB {

	onceSQLx.Do(func() {
		dsn := os.Getenv("DB_URL")

		SQLXConnection, err = sqlx.Connect("mysql", dsn)
		if err != nil {
			utils.FailOnError(err, err.Error())
		}
	})

	utils.LogWithInfo("connected to database", "db")

	return SQLXConnection
}
