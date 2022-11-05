package tests

import (
	"database/sql"
	"log"
	"os"
)

var postgreConnectStr string

func init() {
	postgreConnectStr = os.Getenv("GOLANG_SQL_BENCHMARKS_DSN")
	if postgreConnectStr == "" {
		postgreConnectStr = "root@unix(/var/run/mysqld/mysqld.sock)/golang_sql_benchmarks"
	}
}

func postgreConn() *sql.DB {
	db, err := sql.Open("postgresql", postgreConnectStr)
	if err != nil {
		log.Fatalln(err)
	}
	return db
}
