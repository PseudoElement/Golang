package auth_models

type UserRegister struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Token struct {
	Token string `json:"token"`
}

type UserWithToken struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}

type UserToClient struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
