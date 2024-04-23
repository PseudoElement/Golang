package postgres_main

type TableCreator interface {
	CreateTable() error
}
