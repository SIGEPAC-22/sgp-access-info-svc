package handler

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	kitlog "github.com/go-kit/log"
	"gopkg.in/validator.v2"
	"sgp-access-info-svc/kit/constants"
)

type Middleware func(endpoint endpoint.Endpoint) endpoint.Endpoint

func GetInfoPersonalTransportMiddleware(log kitlog.Logger) Middleware {
	return func(e endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			req := request.(GetInfoPersonalInternalRequest)
			if err := validator.Validate(&req); err != nil {
				log.Log("invalid request", "error", err.Error(), "request", req)
				return GetInfoPersonalInternalResponse{
					Response: constants.ErrBadRequest.Error() + " - " + err.Error(),
					Err:      constants.ErrBadRequest,
				}, nil
			}
			defer log.Log("process finished", "request", req)
			return e(ctx, req)
		}
	}
}
