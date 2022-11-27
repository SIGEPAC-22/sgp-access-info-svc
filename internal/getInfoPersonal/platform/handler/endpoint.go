package handler

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"sgp-access-info-svc/internal/getInfoPersonal"
)

func MakeGetInfoPersonalEndpoint(s getInfoPersonal.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetInfoPersonalInternalRequest)
		resp, err := s.GetInfoPersonalSvc(req.ctx)
		return GetInfoPersonalInternalResponse{
			Response: resp,
			Err:      err,
		}, nil
	}
}

type GetInfoPersonalInternalResponse struct {
	Response interface{}
	Err      error
}

type GetInfoPersonalInternalRequest struct {
	ctx context.Context
}
