package helpers

import (
	"database/sql"
	"fmt"
	"log"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1234"
	dbname   = "go_blog"
)

func DbConn() *sql.DB {

	postgresqlDbInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", postgresqlDbInfo)
	if err != nil {
		log.Fatalf("Connection Lost : %s", err)
	}
	if err != nil {
		panic(err.Error())
	}
	return db
}
