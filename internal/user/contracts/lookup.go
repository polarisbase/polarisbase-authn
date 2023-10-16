package contracts

type LookupRequest struct {
	ID string `json:"id"`
}

type LookupResponse struct {
	User UserDTO `json:"user"`
}
