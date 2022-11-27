package service

import (
	"context"
	kitlog "github.com/go-kit/log"
	"sgp-access-info-svc/internal/getDocumentType"
	"sgp-access-info-svc/kit/constants"
)

type GetDocumentTypeSvc struct {
	repoDB getDocumentType.Repository
	logger kitlog.Logger
}

func NewGetDocumentTypeSvc(repoDB getDocumentType.Repository, logger kitlog.Logger) *GetDocumentTypeSvc {
	return &GetDocumentTypeSvc{repoDB: repoDB, logger: logger}
}

func (g *GetDocumentTypeSvc) GetDocumentTypeSvc(ctx context.Context) ([]getDocumentType.GetDocumentTypeResponse, error) {
	g.logger.Log("Starting subscription", constants.UUID, ctx.Value(constants.UUID))

	resp, err := g.repoDB.GetDocumentTypeRepo(ctx)
	if err != nil {
		g.logger.Log("Error - Information  could not be obtained", constants.UUID, ctx.Value(constants.UUID))
		return []getDocumentType.GetDocumentTypeResponse{}, err
	}
	return resp, nil
}
