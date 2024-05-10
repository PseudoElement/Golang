package chats_queries

type MessageFromDB struct {
	Id        string `json:"id"`
	FromEmail string `json:"from_email"`
	Message   string `json:"message"`
	Date      string `json:"date"`
}

type ChatFromDB struct {
	Messages  []MessageFromDB `json:"messages"`
	Members   []string        `json:"members"`
	Id        string          `json:"id"`
	CreatedAt string          `json:"created_at"`
	UpdatedAt string          `json:"updated_at"`
}
