package service

import (
	"context"
	kitlog "github.com/go-kit/log"
	"sgp-access-info-svc/internal/getOneInfoPersonal"
	"sgp-access-info-svc/kit/constants"
	"strconv"
)

type GetOneInfoPersonalSvc struct {
	repoDB getOneInfoPersonal.Repository
	logger kitlog.Logger
}

func NewGetOneInfoPersonalSvc(repoDB getOneInfoPersonal.Repository, logger kitlog.Logger) *GetOneInfoPersonalSvc {
	return &GetOneInfoPersonalSvc{repoDB: repoDB, logger: logger}
}

func (g *GetOneInfoPersonalSvc) GetOneInfoPersonalSvc(ctx context.Context, id string) (getOneInfoPersonal.GetOneInfoPersonalResponse, error) {
	g.logger.Log("Starting subscription", constants.UUID, ctx.Value(constants.UUID))

	idConverter, _ := strconv.Atoi(id)

	resp, err := g.repoDB.GetOneInfoPersonalRepo(ctx, idConverter)
	if err != nil {
		g.logger.Log("Error - Information  could not be obtained", constants.UUID, ctx.Value(constants.UUID))
		return getOneInfoPersonal.GetOneInfoPersonalResponse{}, err
	}
	return resp, nil
}
