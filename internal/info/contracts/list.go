package contracts

import "github.com/polarisbase/polaris-sdk/v3/services/authn/internal/info/model"

type ListRequest struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type ListResponse struct {
	Infos []model.Info `json:"infos"`
}
