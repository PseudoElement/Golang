package chats_queries

type MessageFromDB struct {
	Id        string `json:"id"`
	FromEmail string `json:"from_email"`
	Message   string `json:"message"`
	Date      string `json:"date"`
}
