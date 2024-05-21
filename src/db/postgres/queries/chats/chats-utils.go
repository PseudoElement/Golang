package chats_queries

import (
	"encoding/json"
	"math"
	"sort"
	"strings"

	errors_module "github.com/pseudoelement/go-server/src/errors"
	string_utils "github.com/pseudoelement/go-server/src/utils/strings"
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

/* page starts from 1 */
func sortMessages(messages []MessageFromDB, sortDir string, page int, limitPerPage int) ([]MessageFromDB, errors_module.ErrorWithStatus) {
	if !string_utils.IsEqual(sortDir, "asc") && !string_utils.IsEqual(sortDir, "desc") {
		return nil, errors_module.DbDefaultError("`sortDir` value should be equal to `asc` or `desc`!")
	}
	sortDir = strings.ToLower(sortDir)

	sort.SliceStable(messages, func(prev, next int) bool {
		if sortDir == "asc" {
			return messages[prev].Date < messages[next].Date
		} else {
			return messages[prev].Date > messages[next].Date
		}
	})

	startIndex := (page - 1) * limitPerPage
	endIndex := math.Min(float64(page*limitPerPage), float64(len(messages)))
	sliced := messages[startIndex:int(endIndex)]

	return sliced, nil
}
