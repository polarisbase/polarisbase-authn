package v1

import (
	"github.com/polarisbase/polarisbase-authn/internal/user/contracts"
	"github.com/polarisbase/polarisbase-shared"
)

func (a *Actions) Lookup(ticket shared.Ticket, request contracts.LookupRequest) (contracts.LookupResponse, error) {

	// Create the response
	response := contracts.LookupResponse{}

	// Lookup in the store
	res, err, ok := a.dep.UserStore.LookupByID(ticket.GetContext(), request.ID)
	if err != nil {
		return response, err
	}

	// Check for ok
	if !ok {
		return response, nil
	}

	// Set the response
	response.User = contracts.UserDTO{
		ID:    res.ID,
		Email: res.Email,
	}

	return response, nil
}
