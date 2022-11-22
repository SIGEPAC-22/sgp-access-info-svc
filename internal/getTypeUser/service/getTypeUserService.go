package service

import (
	"context"
	kitlog "github.com/go-kit/log"
	"sgp-access-info-svc/internal/getTypeUser"
	"sgp-access-info-svc/kit/constants"
)

type GetTypeUserSvc struct {
	repoDB getTypeUser.Repository
	logger kitlog.Logger
}

func NewGetTypeUserSvc(repoDB getTypeUser.Repository, logger kitlog.Logger) *GetTypeUserSvc {
	return &GetTypeUserSvc{repoDB: repoDB, logger: logger}
}

func (g *GetTypeUserSvc) GetTypeUserSvc(ctx context.Context) ([]getTypeUser.GetTypeUserResponse, error) {
	g.logger.Log("Starting subscription", constants.UUID, ctx.Value(constants.UUID))

	resp, err := g.repoDB.GetTypeUserRepo(ctx)
	if err != nil {
		g.logger.Log("Error - Information  could not be obtained", constants.UUID, ctx.Value(constants.UUID))
		return []getTypeUser.GetTypeUserResponse{}, err
	}
	return resp, nil
}
