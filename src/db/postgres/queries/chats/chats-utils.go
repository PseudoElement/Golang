package chats_queries

import (
	"encoding/json"

	errors_module "github.com/pseudoelement/go-server/src/errors"
)

func parseStringArrayIntoJsonArray[T any](arr []string) ([]T, errors_module.ErrorWithStatus) {
	var typedArr []T
	for _, str := range arr {
		var el T
		err := json.Unmarshal([]byte(str), &el)
		if err != nil {
			return nil, errors_module.DbDefaultError(err.Error())
		}
		typedArr = append(typedArr, el)
	}

	return typedArr, nil
}
