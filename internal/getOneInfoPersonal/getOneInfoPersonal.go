package getOneInfoPersonal

import "context"

type Repository interface {
	GetOneInfoPersonalRepo(ctx context.Context, DocumentNumber string) (GetOneInfoPersonalResponse, error)
}

type Service interface {
	GetOneInfoPersonalSvc(ctx context.Context, DocumentNumber string) (GetOneInfoPersonalResponse, error)
}

type GetOneInfoPersonalRequest struct {
	DocumentNumber string `json:"documentNumber"`
}

type GetOneInfoPersonalResponse struct {
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
