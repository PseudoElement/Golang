package chats_queries

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
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
	membersValue := fmt.Sprintf("ARRAY[%v]", membersJoined)
	messagesValue := "ARRAY[]"
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
	err := row.Scan(&chat.Id, &chat.Messages, &chat.Members, &chat.CreatedAt, &chat.UpdatedAt)

	return postgres_main.HandleQueryRowErrors(chat, err)
}

func (cq *ChatsQueries) GetChatByMembers(members ...string) (ChatFromDB, errors_module.ErrorWithStatus) {
	membersJoined := strings.Join(members, ",")
	query := `
		SELECT * FROM chats 
		WHERE members && '{$1}';
	`
	row := cq.db.QueryRow(query, membersJoined)
	var chat ChatFromDB
	err := row.Scan(&chat.Id, &chat.Messages, &chat.Members, &chat.CreatedAt, &chat.UpdatedAt)

	return postgres_main.HandleQueryRowErrors(chat, err)
}

func (cq *ChatsQueries) AddMessage(chatId string, from_email string, message string) errors_module.ErrorWithStatus {
	messageId := uuid.New().String()
	newValue := fmt.Sprintf(`
        {
         "from_email": %v,
         "message": %v,
         "date": %v,
         "id": %v
        }
    `, from_email, message, time.Now(), messageId)
	query := `
        UPDATE chats 
        SET messages = ARRAY_APPEND(messages, $1)
        WHERE id = $2;
    `

	r, err := cq.db.Exec(query, newValue, chatId)

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

func (cq *ChatsQueries) GetChatMessages(chatId string) ([]MessageFromDB, errors_module.ErrorWithStatus) {
	query := `
		SELECT messages FROM chats 
		WHERE id = $1;
	`
	row := cq.db.QueryRow(query, chatId)
	var messages []MessageFromDB
	err := row.Scan(&messages)

	return postgres_main.HandleQueryRowErrors(messages, err)
}

func (cq *ChatsQueries) GetAllChatsOfUser(email string) ([]ChatFromDB, errors_module.ErrorWithStatus) {
	query := `
		SELECT * FROM chats
		WHERE members && '{$1}'
	`
	rows, queryErr := cq.db.Query(query, email)
	if queryErr != nil {
		return nil, errors_module.DbDefaultError(queryErr.Error())
	}
	defer rows.Close()

	var chats []ChatFromDB
	for rows.Next() {
		var chat ChatFromDB
		err := rows.Scan(&chat.Id, &chat.Messages, &chat.Members, &chat.CreatedAt, &chat.UpdatedAt)
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
