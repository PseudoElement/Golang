package cards

type NewCard struct {
	Author string `json:"author"`
	Info   string `json:"info"`
}

type CardUpdate struct {
	NewCard
	Id string `json:"id"`
}

type CardDelete struct {
	Id string `json:"id"`
}

type CardGet struct {
	Id string `json:"id"`
}

type CardToClient struct {
	Author    string `json:"author"`
	Info      string `json:"info"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Id        string `json:"id"`
}
