package handler

type LoginRequest struct {
	Email    string
	Password string
}

type RegisterRequest struct {
	Fullname  string
	Username  string
	Handphone string
	Email     string
	Password  string
}
