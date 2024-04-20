package postgres_main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "postgres"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

var DB *sql.DB

func Init() {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", 
		host, 
		port, 
		user, 
		password, 
		dbname)
	DB, _ = sql.Open("postgres", psqlInfo)

	// createDbIfNotExists()
	createTestTableIfNotExists()
	defer DB.Close()

	err := DB.Ping()
	if err != nil{
		panic(err)
	}
}