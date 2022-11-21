package getInfoPersonal

import "context"

type Repository interface {
	GetInfoPersonalRepo(ctx context.Context) ([]GetInfoPersonalResponse, error)
}

type Service interface {
	GetInfoPersonalSvc(ctx context.Context) ([]GetInfoPersonalResponse, error)
}

type GetInfoPersonalResponse struct {
	Id                  int    `json:"id"`
	FirstName           string `json:"firstName"`
	SecondName          string `json:"secondName"`
	LastName            string `json:"lastName"`
	SecondLastName      string `json:"secondLastName"`
	Sex                 string `json:"sex"`
	DateOfBirth         string `json:"dateOfBirth"`
	TypeDocument        string `json:"typeDocument"`
	DocumentNumber      string `json:"documentNumber"`
	User                string `json:"user"`
	TypeUser            string `json:"typeUser"`
	DateCreationAccount string `json:"dateCreationAccount"`
	StateName           string `json:"stateName"`
}
