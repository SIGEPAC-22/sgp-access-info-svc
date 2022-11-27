package handler

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"sgp-access-info-svc/internal/getOneInfoPersonal"
)

func MakeGetOneInfoPersonalEndpoint(s getOneInfoPersonal.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetOneInfoPersonalInternalRequest)
		resp, err := s.GetOneInfoPersonalSvc(req.ctx, req.Id)
		return GetOneInfoPersonalInternalResponse{
			Response: resp,
			Err:      err,
		}, nil
	}
}

type GetOneInfoPersonalInternalResponse struct {
	Response interface{}
	Err      error
}

type GetOneInfoPersonalInternalRequest struct {
	Id  string `json:"id"`
	ctx context.Context
}
