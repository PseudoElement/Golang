package postgres_main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type PostgresDB struct {
	Host     string
	Port     int32
	User     string
	Password string
	Name     string
	db       *sql.DB
}

func GetInstance() *PostgresDB {
	return &PostgresDB{
		Host:     "postgres",
		Port:     5432,
		User:     "postgres",
		Password: "postgres",
		Name:     "postgres",
	}
}

func (pg *PostgresDB) GetDB() *sql.DB {
	return pg.db
}

func (pg *PostgresDB) Init() {
	db, err := pg.connectDB()
	if err != nil {
		panic(err)
	}
	pg.db = db

	createTestTableIfNotExists(db)

	err = db.Ping()
	if err != nil {
		panic(err)
	}
}

func (pg *PostgresDB) connectDB() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		pg.Host,
		pg.Port,
		pg.User,
		pg.Password,
		pg.Name)

	return sql.Open("postgres", psqlInfo)
}
