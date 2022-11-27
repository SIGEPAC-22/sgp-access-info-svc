package handler

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"sgp-access-info-svc/internal/getPersonalSex"
)

func MakeGetPersonalSexEndpoints(c getPersonalSex.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetPersonalSexInternalRequest)
		resp, err := c.GetPersonalSexSvc(req.ctx)
		return GetPersonalSexInternalResponse{
			Response: resp,
			Err:      err,
		}, nil
	}
}

type GetPersonalSexInternalResponse struct {
	Response interface{}
	Err      error
}

type GetPersonalSexInternalRequest struct {
	ctx context.Context
}
