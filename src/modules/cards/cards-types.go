package cards

type NewCard struct {
	Author string `json:"author"`
	Info   string `json:"info"`
}

type CardUpdate struct {
	Author string `json:"author"`
	Info   string `json:"info"`
	Id     string `json:"id"`
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
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	Id        string `json:"id"`
}

type CardActionSuccess struct {
	Message string `json:"message"`
}
