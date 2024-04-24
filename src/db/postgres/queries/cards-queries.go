package queries

import (
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	errors_module "github.com/pseudoelement/go-server/src/errors"
)

type CardsQueries struct {
	db *sql.DB
}

func NewCardsQueries(db *sql.DB) *CardsQueries {
	return &CardsQueries{
		db: db,
	}
}

func (cq *CardsQueries) CreateTable() error {
	_, err := cq.db.Exec(`
		CREATE TABLE IF NOT EXISTS cards (
			id varchar(255) NOT NULL PRIMARY KEY, 
			info varchar(255) NOT NULL, 
			author varchar(255) NOT NULL, 
			created_at timestamp NOT NULL,
			updated_at timestamp NOT NULL
		);
	`)
	if err != nil {
		return errors_module.DbDefaultError(err.Error())
	}

	return nil
}

func (cq *CardsQueries) AddCard(author string, info string) errors_module.ErrorWithStatus {
	id := uuid.New().String()
	_, err := cq.db.Exec(`
		INSERT INTO cards (id, info, author, created_at, updated_at) 
		VALUES ($1, $2, $3, $4, $5);
	`, id, author, info, time.Now(), time.Now())
	if err != nil {
		return errors_module.DbDefaultError(err.Error())
	}

	return nil
}

func (cq *CardsQueries) UpdateCard(id string, newInfo string, newAuthor string) errors_module.ErrorWithStatus {
	var query string

	if newAuthor != "" {
		query = `UPDATE cards SET info = $1, author = $2 WHERE id = $3;`
		_, err := cq.db.Exec(query, newInfo, newAuthor, id)
		if err != nil {
			return errors_module.DbDefaultError(err.Error())
		}
	} else {
		query = `UPDATE cards SET info = $1 WHERE id = $2;`
		_, err := cq.db.Exec(query, newInfo, id)
		if err != nil {
			return errors_module.DbDefaultError(err.Error())
		}
	}

	return nil
}

func (cq *CardsQueries) DeleteCard(id string) errors_module.ErrorWithStatus {
	_, err := cq.db.Exec(`DELETE FROM cards WHERE id = $1;`, id)
	if err != nil {
		return errors_module.DbDefaultError(err.Error())
	}

	return nil
}

func (cq *CardsQueries) GetCard(id string) (CardFromDB, errors_module.ErrorWithStatus) {
	row := cq.db.QueryRow(`SELECT * FROM cards WHERE id = $1;`, id)
	var card CardFromDB

	err := row.Scan(&card.Id, &card.Info, &card.Author, &card.CreatedAt, &card.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return CardFromDB{}, errors_module.DbDefaultError("Card not found!")
		}
		return CardFromDB{}, errors_module.DbDefaultError(err.Error())
	}

	return card, nil
}

func (cq *CardsQueries) GetAllSortedCard(sortBy string, sortDir string, page int, limitPerPage int) ([]CardFromDB, errors_module.ErrorWithStatus) {
	query := `
		SELECT * FROM cards 
		ORDER BY ` + sortBy + ` ` + sortDir + `
		LIMIT $1 OFFSET $2;
	`
	offset := (page - 1) * limitPerPage

	rows, queryErr := cq.db.Query(query, limitPerPage, offset)
	if queryErr != nil {
		return nil, errors_module.DbDefaultError(queryErr.Error())
	}
	defer rows.Close()

	var cards []CardFromDB
	for rows.Next() {
		var card CardFromDB
		if err := rows.Scan(&card.Id, &card.Info, &card.Author, &card.CreatedAt, &card.UpdatedAt); err != nil {
			return nil, errors_module.DbDefaultError(err.Error())
		}
		cards = append(cards, card)
	}

	if err := rows.Err(); err != nil {
		return nil, errors_module.DbDefaultError(err.Error())
	}

	return cards, nil
}
