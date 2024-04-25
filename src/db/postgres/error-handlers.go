package postgres_main

import (
	"database/sql"
	"errors"

	errors_module "github.com/pseudoelement/go-server/src/errors"
)

func HandleExecErrors(r sql.Result, err error) errors_module.ErrorWithStatus {
	if err != nil {
		return errors_module.DbDefaultError(err.Error())
	}
	count, _ := r.RowsAffected()
	if count < 1 {
		return errors_module.DbDefaultError("Value by provided data was not found in database!")
	}

	return nil
}

func HandleQueryRowErrors[T any](value T, err error) (T, errors_module.ErrorWithStatus) {
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return value, errors_module.DbDefaultError("Card not found!")
		}
		return value, errors_module.DbDefaultError(err.Error())
	}

	return value, nil
}
