package handler

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"sgp-access-info-svc/internal/getTypeUser"
)

func MakeGetTypeUserEndpoints(c getTypeUser.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetTypeUserInternalRequest)
		resp, err := c.GetTypeUserSvc(req.ctx)
		return GetTypeUserInternalResponse{
			Response: resp,
			Err:      err,
		}, nil
	}
}

type GetTypeUserInternalResponse struct {
	Response interface{}
	Err      error
}

type GetTypeUserInternalRequest struct {
	ctx context.Context
}
