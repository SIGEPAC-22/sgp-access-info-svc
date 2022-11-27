package getPersonalSex

import "context"

type Repository interface {
	GetPersonalSexRepo(ctx context.Context) ([]GetPersonalSexResponse, error)
}

type Service interface {
	GetPersonalSexSvc(ctx context.Context) ([]GetPersonalSexResponse, error)
}

type GetPersonalSexResponse struct {
	Id      int    `json:"id"`
	NameSex string `json:"nameSex"`
}
