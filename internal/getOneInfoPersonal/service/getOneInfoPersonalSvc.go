package service

import (
	"context"
	kitlog "github.com/go-kit/log"
	"sgp-access-info-svc/internal/getOneInfoPersonal"
	"sgp-access-info-svc/kit/constants"
)

type GetOneInfoPersonalSvc struct {
	repoDB getOneInfoPersonal.Repository
	logger kitlog.Logger
}

func NewGetOneInfoPersonalSvc(repoDB getOneInfoPersonal.Repository, logger kitlog.Logger) *GetOneInfoPersonalSvc {
	return &GetOneInfoPersonalSvc{repoDB: repoDB, logger: logger}
}

func (g *GetOneInfoPersonalSvc) GetOneInfoPersonalSvc(ctx context.Context, DocumentNumber string) (getOneInfoPersonal.GetOneInfoPersonalResponse, error) {
	g.logger.Log("Starting subscription", constants.UUID, ctx.Value(constants.UUID))

	resp, err := g.repoDB.GetOneInfoPersonalRepo(ctx, DocumentNumber)
	if err != nil {
		g.logger.Log("Error - Information  could not be obtained", constants.UUID, ctx.Value(constants.UUID))
		return getOneInfoPersonal.GetOneInfoPersonalResponse{}, err
	}
	return resp, nil
}
