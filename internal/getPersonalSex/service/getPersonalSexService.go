package service

import (
	"context"
	kitlog "github.com/go-kit/log"
	"sgp-access-info-svc/internal/getPersonalSex"
	"sgp-access-info-svc/kit/constants"
)

type GetPersonalSexSvc struct {
	repoDB getPersonalSex.Repository
	logger kitlog.Logger
}

func NewGetPersonalSexSvc(repoDB getPersonalSex.Repository, logger kitlog.Logger) *GetPersonalSexSvc {
	return &GetPersonalSexSvc{repoDB: repoDB, logger: logger}
}

func (g *GetPersonalSexSvc) GetPersonalSexSvc(ctx context.Context) ([]getPersonalSex.GetPersonalSexResponse, error) {
	g.logger.Log("Starting subscription", constants.UUID, ctx.Value(constants.UUID))

	resp, err := g.repoDB.GetPersonalSexRepo(ctx)
	if err != nil {
		g.logger.Log("Error - Information  could not be obtained", constants.UUID, ctx.Value(constants.UUID))
		return []getPersonalSex.GetPersonalSexResponse{}, err
	}
	return resp, nil
}
