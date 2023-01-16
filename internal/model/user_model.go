package model

type RegisterUserPayload struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Role  string `json:"role"`
}

type RegisterUserResponse struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Role     string `json:"role"`
	Password string `json:"password"`
}

type GetUserPayload struct {
	Phone    string `json:"phone"`
}

type GetUserResponse struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Role     string `json:"role"`
	Password string `json:"password"`
}

type LoginPayload struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
