package postgres_main

import errors_module "github.com/pseudoelement/go-server/src/errors"

type TableCreator interface {
	CreateTable() errors_module.ErrorWithStatus
}
