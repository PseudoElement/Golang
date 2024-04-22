package postgres_main

import (
	"database/sql"
	"fmt"
)

// func createDbIfNotExists(db *sql.DB, dbName string) {
// 	_, err := db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %v", dbName))
// 	if err != nil {
// 		msg := fmt.Sprintf("Error creating %v database - %v", dbName, err)
// 		panic(msg)
// 	}
// }

func createTestTableIfNotExists(db *sql.DB) {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS test_table(last_name varchar(255), first_name varchar(255));")
	if err != nil {
		msg := fmt.Sprintf("Error creating test_table - %v", err)
		panic(msg)
	}
}
