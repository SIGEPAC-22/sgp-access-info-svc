package getTypeUser

import "context"

type Repository interface {
	GetTypeUserRepo(ctx context.Context) ([]GetTypeUserResponse, error)
}

type Service interface {
	GetTypeUserSvc(ctx context.Context) ([]GetTypeUserResponse, error)
}

type GetTypeUserResponse struct {
	Id       int    `json:"id"`
	NameType string `json:"nameType"`
}
