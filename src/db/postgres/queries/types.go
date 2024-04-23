package queries

type CardFromDB struct {
	Author    string `json:"author"`
	Info      string `json:"info"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Id        string `json:"id"`
}
