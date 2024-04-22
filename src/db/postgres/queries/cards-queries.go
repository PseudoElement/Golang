package queries

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type CardsQueries struct {
	db *sql.DB
}

func GetInstance(db *sql.DB) *CardsQueries {
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
		return err
	}

	return nil
}

func (cq *CardsQueries) AddCard(author string, info string) error {
	id := uuid.New().String()
	card, err := cq.db.Exec(`
		INSERT INTO cards (id, info, author, created_at, updated_at) 
		VALUES ($1, $2, $3, $4, $5);
	`, id, author, info, time.Now(), time.Now())
	if err != nil {
		return err
	}
	fmt.Println("NEW CARD - ", card)

	return nil
}

func (cq *CardsQueries) UpdateCard(id int, newInfo string, newAuthor string) error {
	var query string

	if newAuthor != "" {
		query = `UPDATE cards SET info = $1, author = $2 WHERE id = $3;`
		_, err := cq.db.Exec(query, newInfo, newAuthor, id)
		if err != nil {
			return err
		}
	} else {
		query = `UPDATE cards SET info = $1 WHERE id = $2;`
		upd, err := cq.db.Exec(query, newInfo, id)
		if err != nil {
			return err
		}
		fmt.Println("UPDATED CARD - ", upd)
	}

	return nil
}

func (cq *CardsQueries) DeleteCard(id int) error {
	del, err := cq.db.Exec(`DELETE FROM cards WHERE id = $1;`, id)
	if err != nil {
		return err
	}
	fmt.Println("DELETED CARD - ", del)

	return nil
}
