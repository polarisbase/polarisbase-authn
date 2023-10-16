package v1

import (
	"fmt"
	"github.com/polarisbase/polaris-sdk/v3/services/authn/internal/user/contracts"
	"github.com/polarisbase/polaris-sdk/v3/services/authn/internal/user/model"
	"github.com/polarisbase/polaris-sdk/v3/shared"
	"golang.org/x/crypto/bcrypt"
)

func (a *Actions) Create(ticket shared.Ticket, request contracts.CreateRequest) (contracts.CreateResponse, error) {
	// Create the response
	response := contracts.CreateResponse{}

	// Validate email
	err, ok := a.validateEmail(request.Email)
	// Check for errors
	if err != nil {
		return response, err
	}
	// Check for ok
	if !ok {
		return response, nil
	}

	// Hash the password
	hashedPassword, err := a.hashPassword(request.Password)
	if err != nil {
		return response, err
	}

	// Create in the store
	res, err, ok := a.dep.UserStore.CreateUser(ticket.GetContext(), model.User{
		Email:        request.Email,
		PasswordHash: hashedPassword,
	})

	// Check for errors
	if err != nil {
		return response, err
	}

	// Check for ok
	if !ok {
		return response, nil
	}

	// Set the response
	response.ID = res.ID

	// Return the response
	return response, nil
}

func (a *Actions) hashPassword(password string) (passwordHash string, err error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	if err != nil {
		return "", err
	}
	passwordHash = string(hashed)
	return passwordHash, nil
}

func (a *Actions) validateEmail(email string) (err error, ok bool) {
	// Check for empty
	if email == "" {
		return fmt.Errorf("email is required"), false
	}

	// Check for length
	if len(email) > 255 {
		return fmt.Errorf("email is too long"), false
	}

	// Check if email is in use
	err, ok = a.dep.UserStore.CheckIfEmailIsAlreadyInUse(email)

	// Check for errors
	if err != nil {
		return err, false
	}

	// Check for ok
	if !ok {
		return fmt.Errorf("email is already in use"), false
	}

	// Return nil
	return nil, true
}
