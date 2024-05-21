package cards

import (
	"strconv"
	"strings"

	cards_queries "github.com/pseudoelement/go-server/src/db/postgres/queries/cards"
	errors_module "github.com/pseudoelement/go-server/src/errors"
	slice_utils "github.com/pseudoelement/go-server/src/utils/slices"
)

func (m *CardsModule) getSortedCards(params map[string]string) ([]cards_queries.CardFromDB, errors_module.ErrorWithStatus) {
	page, _ := strconv.Atoi(params["page"])
	limitPerPage, _ := strconv.Atoi(params["limitPerPage"])

	if err := m.checkGetSortedCardsQueryValues(params); err != nil {
		return nil, err
	}

	cards, err := m.cardsQueries.GetAllSortedCard(
		params["sortBy"],
		params["sortDir"],
		page,
		limitPerPage)
	if err != nil {
		return nil, err
	}

	return cards, nil
}

func (m *CardsModule) checkGetSortedCardsQueryValues(params map[string]string) errors_module.ErrorWithStatus {
	sortByValues := []string{"created_at", "updated_at", "author", "info"}
	sortDirValues := []string{"asc", "desc"}
	sortByToLower := strings.ToLower(params["sortBy"])
	sortDirToLower := strings.ToLower(params["sortDir"])

	if !slice_utils.Contains(sortByValues, sortByToLower) {
		return errors_module.IncorrectQueryParamValue("sortBy")
	}
	if !slice_utils.Contains(sortDirValues, sortDirToLower) {
		return errors_module.IncorrectQueryParamValue("sortDir")
	}

	return nil
}
