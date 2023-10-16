package contracts

type CreateRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r CreateRequest) Validate() error {
	if r.Email == "" {
		return ErrEmailIsRequired
	}
	if r.Password == "" {
		return ErrPasswordIsRequired
	}
	return nil
}

type CreateResponse struct {
	ID string `json:"id"`
}
