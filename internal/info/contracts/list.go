package contracts

import "github.com/polarisbase/polarisbase-authn/internal/info/model"

type ListRequest struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type ListResponse struct {
	Infos []model.Info `json:"infos"`
}
