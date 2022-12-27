package auth

type AuthLoginInputDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
