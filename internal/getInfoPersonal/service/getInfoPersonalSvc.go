package service

import (
	"context"
	kitlog "github.com/go-kit/log"
	"sgp-access-info-svc/internal/getInfoPersonal"
	"sgp-access-info-svc/kit/constants"
)

type GetInfoPersonalSvc struct {
	repoDB getInfoPersonal.Repository
	logger kitlog.Logger
}

func NewGetInfoPersonalSvc(repoDB getInfoPersonal.Repository, logger kitlog.Logger) *GetInfoPersonalSvc {
	return &GetInfoPersonalSvc{repoDB: repoDB, logger: logger}
}

func (g *GetInfoPersonalSvc) GetInfoPersonalSvc(ctx context.Context) ([]getInfoPersonal.GetInfoPersonalResponse, error) {
	g.logger.Log("Starting subscription", constants.UUID, ctx.Value(constants.UUID))

	resp, err := g.repoDB.GetInfoPersonalRepo(ctx)
	if err != nil {
		g.logger.Log("Error - Information  could not be obtained", constants.UUID, ctx.Value(constants.UUID))
		return []getInfoPersonal.GetInfoPersonalResponse{}, err
	}
	return resp, nil
}
