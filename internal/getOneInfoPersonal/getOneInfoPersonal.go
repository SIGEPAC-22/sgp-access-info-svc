package getOneInfoPersonal

import "context"

type Repository interface {
	GetOneInfoPersonalRepo(ctx context.Context, id int) (GetOneInfoPersonalResponse, error)
}

type Service interface {
	GetOneInfoPersonalSvc(ctx context.Context, id string) (GetOneInfoPersonalResponse, error)
}

type GetOneInfoPersonalRequest struct {
	Id string `json:"id"`
}

type GetOneInfoPersonalResponse struct {
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
	Password            string `json:"password"`
	TypeUser            string `json:"typeUser"`
	DateCreationAccount string `json:"dateCreationAccount"`
	StateName           string `json:"stateName"`
}
