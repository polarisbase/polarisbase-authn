package v1

import (
	"github.com/polarisbase/polarisbase-authn/internal/info/contracts"
	"github.com/polarisbase/polarisbase-shared"
)

func (a *Actions) List(ticket shared.Ticket, request contracts.ListRequest) (contracts.ListResponse, error) {
	// Create the response
	response := contracts.ListResponse{}

	// Create in the store
	res, err, ok := a.dep.InfoStore.List(ticket.GetContext(), request.Limit, request.Offset)
	if err != nil {
		return response, err
	}
	if !ok {
		return response, nil
	}

	// Set the response
	response.Infos = res

	// Return the response
	return response, nil
}
