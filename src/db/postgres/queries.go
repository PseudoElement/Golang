package postgres_main

import "fmt"

func createDbIfNotExists() {
	_, err := DB.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s;", dbname))
	if err != nil {
		msg := fmt.Sprintf("Error creating %v database - %v", dbname, err)
		panic(msg)
	}
}

func createTestTableIfNotExists() {
	_, err := DB.Exec("CREATE TABLE IF NOT EXISTS test_table(LastName varchar(255), FirstName varchar(255));")
	if err != nil {
		msg := fmt.Sprintf("Error creating test_table - %v", err)
		panic(msg)
	}
}