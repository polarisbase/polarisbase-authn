package v1

import (
	"github.com/polarisbase/polaris-sdk/v3/services/authn/internal/user/contracts"
	"github.com/polarisbase/polaris-sdk/v3/shared"
	"golang.org/x/crypto/bcrypt"
)

func (a *Actions) Login(ticket shared.Ticket, request contracts.LoginRequest) (contracts.LoginResponse, error) {

	// Create the response
	response := contracts.LoginResponse{}

	// Lookup in the store
	res, err, ok := a.dep.UserStore.LookupByEmail(ticket.GetContext(), request.Email)
	if err != nil {
		return response, err
	}

	// Check for ok
	if !ok {
		return response, nil
	}

	// Check the password
	err = a.checkPassword(request.Password, res.PasswordHash)
	if err != nil {
		return response, err
	}

	// Set the response
	response.User = contracts.UserDTO{
		ID:    res.ID,
		Email: res.Email,
	}

	return response, nil
}

func (a *Actions) checkPassword(password string, passwordHash string) error {
	return bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
}
