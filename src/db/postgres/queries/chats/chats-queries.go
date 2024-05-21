package chats_queries

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	postgres_main "github.com/pseudoelement/go-server/src/db/postgres"
	errors_module "github.com/pseudoelement/go-server/src/errors"
)

type ChatsQueries struct {
	db *sql.DB
}

func NewChatsQueries(db *sql.DB) *ChatsQueries {
	return &ChatsQueries{
		db: db,
	}
}

func (cq *ChatsQueries) CreateTable() errors_module.ErrorWithStatus {
	_, err := cq.db.Exec(`
		CREATE TABLE IF NOT EXISTS chats (
			id varchar(255) NOT NULL PRIMARY KEY, 
			messages jsonb[] NOT NULL, 
			members text[] NOT NULL, 
			created_at timestamp NOT NULL,
            updated_at timestamp NOT NULL
		);
	`)
	if err != nil {
		return errors_module.DbDefaultError(err.Error())
	}

	return nil
}

func (cq *ChatsQueries) DeleteMemberFromChat(email string, chatId string) errors_module.ErrorWithStatus {
	query := `
		UPDATE chats
		SET members = array_remove(members, $1)
		WHERE id = $2;
	`
	r, err := cq.db.Exec(query, email, chatId)

	return postgres_main.HandleExecErrors(r, err)
}

func (cq *ChatsQueries) AddMemberInChat(email string, chatId string) errors_module.ErrorWithStatus {
	query := `
		UPDATE chats
		SET members = array_append(members, $1)
		WHERE id = $2;
	`
	r, err := cq.db.Exec(query, email, chatId)

	return postgres_main.HandleExecErrors(r, err)
}

func (cq *ChatsQueries) CreateChat(members ...string) (string, errors_module.ErrorWithStatus) {
	id := uuid.New().String()
	membersJoined := strings.Join(members, ",")
	membersValue := fmt.Sprintf("{%v}", membersJoined)
	messagesValue := "{}"
	created_at := time.Now()
	updated_at := created_at

	query := `
        INSERT INTO chats (id, members, messages, created_at, updated_at) 
        VALUES ($1, $2, $3, $4, $5);
    `
	r, err := cq.db.Exec(
		query,
		id,
		membersValue,
		messagesValue,
		created_at,
		updated_at)

	execErr := postgres_main.HandleExecErrors(r, err)
	if execErr != nil {
		return "", execErr
	}

	return id, nil
}

func (cq *ChatsQueries) DeleteChatById(chatId string) errors_module.ErrorWithStatus {
	r, err := cq.db.Exec(`DELETE FROM chats WHERE id = $1;`, chatId)

	return postgres_main.HandleExecErrors(r, err)
}

func (cq *ChatsQueries) GetChatById(chatId string) (ChatFromDB, errors_module.ErrorWithStatus) {
	query := `
		SELECT * FROM chats
		WHERE id = $1;
	`
	row := cq.db.QueryRow(query, chatId)

	var chat ChatFromDB
	var messagesStrings []string
	err := row.Scan(&chat.Id, pq.Array(&messagesStrings), pq.Array(&chat.Members), &chat.CreatedAt, &chat.UpdatedAt)
	if err != nil {
		return chat, errors_module.DbDefaultError(err.Error())
	}

	messages, err := parseStringArrayIntoJsonArray[MessageFromDB](messagesStrings)
	if err != nil {
		return chat, errors_module.DbDefaultError(err.Error())
	}
	chat.Messages = messages

	return postgres_main.HandleQueryRowErrors(chat, err)
}

func (cq *ChatsQueries) GetChatByMembers(members ...string) (ChatFromDB, errors_module.ErrorWithStatus) {
	// @> checks db-array on containing all passed members
	// && checks db-array on containig at least one passed member
	query := `
		SELECT * FROM chats 
		WHERE members @> $1;
	`
	row := cq.db.QueryRow(query, pq.Array(members))

	var chat ChatFromDB
	err := row.Scan(&chat.Id, pq.Array(&chat.Messages), pq.Array(&chat.Members), &chat.CreatedAt, &chat.UpdatedAt)

	return postgres_main.HandleQueryRowErrors(chat, err)
}

func (cq *ChatsQueries) AddMessage(chatId string, from_email string, message string) errors_module.ErrorWithStatus {
	messageId := uuid.New().String()
	newMessageBytes, err := json.Marshal(MessageFromDB{
		Id:        messageId,
		FromEmail: from_email,
		Message:   message,
		Date:      time.Now().String(),
	})
	if err != nil {
		return errors_module.DbDefaultError(err.Error())
	}

	query := `
        UPDATE chats 
        SET messages = ARRAY_APPEND(messages, $1)
        WHERE id = $2;
    `

	r, err := cq.db.Exec(query, newMessageBytes, chatId)

	return postgres_main.HandleExecErrors(r, err)
}

func (cq *ChatsQueries) GetMessageById(chatId string, messageId string) (MessageFromDB, errors_module.ErrorWithStatus) {
	query := `
        SELECT msg->>'message' AS message,
               msg->>'from_email' AS from_email,
               msg->>'date' AS date,
               msg->>'id' AS id
        FROM chats,
        jsonb_array_elements(messages) AS msg
        WHERE id = $1
        AND msg->>'id' = $2;
    `
	row := cq.db.QueryRow(query, chatId, messageId)

	var message MessageFromDB
	err := row.Scan(&message.FromEmail, &message.Message, &message.Date, &message.Id)

	return postgres_main.HandleQueryRowErrors(message, err)
}

func (cq *ChatsQueries) DeleteMessage(chatId string, messageId string) errors_module.ErrorWithStatus {
	message, err := cq.GetMessageById(chatId, messageId)
	if err != nil {
		return errors_module.DbDefaultError(err.Error())
	}

	query := `
        SELECT array_remove(messages, $1) AS reduced_array
        FROM chats 
        WHERE id = $2;
    `
	r, execErr := cq.db.Exec(query, chatId, message)

	return postgres_main.HandleExecErrors(r, execErr)
}

func (cq *ChatsQueries) GetChatMessages(chatId string, sortDir string, page int, limitPerPage int) ([]MessageFromDB, errors_module.ErrorWithStatus) {
	query := `
		SELECT messages FROM chats 
		WHERE id = $1;
	`
	row := cq.db.QueryRow(query, chatId)

	var messagesStrings []string
	err := row.Scan(pq.Array(&messagesStrings))
	if err != nil {
		return nil, errors_module.DbDefaultError(err.Error())
	}

	messages, err := parseStringArrayIntoJsonArray[MessageFromDB](messagesStrings)
	if err != nil {
		return nil, errors_module.DbDefaultError(err.Error())
	}

	messages, err = sortMessages(messages, sortDir, page, limitPerPage)

	return postgres_main.HandleQueryRowErrors(messages, err)
}

func (cq *ChatsQueries) GetAllChatsOfUser(email string) ([]ChatFromDB, errors_module.ErrorWithStatus) {
	// @> checks db-array on containing all passed members
	// && checks db-array on containig at least one passed member
	query := `
		SELECT * FROM chats
		WHERE members && $1;
	`
	rows, queryErr := cq.db.Query(query, pq.Array([]string{email}))
	if queryErr != nil {
		return nil, errors_module.DbDefaultError(queryErr.Error())
	}
	defer rows.Close()

	var chats []ChatFromDB
	for rows.Next() {
		var chat ChatFromDB
		err := rows.Scan(&chat.Id, pq.Array(&chat.Messages), pq.Array(&chat.Members), &chat.CreatedAt, &chat.UpdatedAt)
		if err != nil {
			return nil, errors_module.DbDefaultError(err.Error())
		}
		chats = append(chats, chat)
	}

	if err := rows.Err(); err != nil {
		return nil, errors_module.DbDefaultError(err.Error())
	}

	return chats, nil
}
