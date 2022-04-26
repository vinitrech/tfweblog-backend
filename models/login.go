package models

type Login struct {
	Email string `json:"email"`
	Senha string `json:"senha"`
}
type GoogleLogin struct {
	Email string `json:"email"`
}