package contracts

import "fmt"

var (
	ErrEmailIsRequired    = fmt.Errorf("email is required")
	ErrPasswordIsRequired = fmt.Errorf("password is required")
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r LoginRequest) Validate() error {
	if r.Email == "" {
		return ErrEmailIsRequired
	}
	if r.Password == "" {
		return ErrPasswordIsRequired
	}
	return nil
}

type LoginResponse struct {
	User UserDTO `json:"user"`
}
