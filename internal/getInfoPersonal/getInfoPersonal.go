package getInfoPersonal

import "context"

type Repository interface {
	GetInfoPersonalRepo(ctx context.Context, DocumentNumber string) (GetInfoPersonalResponse, error)
}

type Service interface {
	GetInfoPersonalSvc(ctx context.Context, DocumentNumber string) (GetInfoPersonalResponse, error)
}

type GetInfoPersonalRequest struct {
	DocumentNumber string `json:"documentNumber"`
}

type GetInfoPersonalResponse struct {
	FirstName       string `json:"firstName"`
	SecondName      string `json:"secondName"`
	LastName        string `json:"lastName"`
	MothersLastName string `json:"mothersLastName"`
	Sex             string `json:"sex"`
	DateOfBirth     string `json:"dateOfBirth"`
	DocumentNumber  string `json:"documentNumber"`
	User            string `json:"user"`
}
