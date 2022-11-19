package handler

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"sgp-access-info-svc/internal/getDocumentType"
)

func MakeGetDocumentTypeEndpoints(c getDocumentType.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetDocumentTypeInternalRequest)
		resp, err := c.GetDocumentTypeSvc(req.ctx)
		return GetDocumentTypeInternalResponse{
			Response: resp,
			Err:      err,
		}, nil
	}
}

type GetDocumentTypeInternalResponse struct {
	Response interface{}
	Err      error
}

type GetDocumentTypeInternalRequest struct {
	ctx context.Context
}
