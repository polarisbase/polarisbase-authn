package model

type User struct {
	ID           string `json:"id" db:"id"`
	Email        string `json:"email" db:"email" unique:"true"`
	PasswordHash string `json:"password_hash" db:"password_hash"`
}
