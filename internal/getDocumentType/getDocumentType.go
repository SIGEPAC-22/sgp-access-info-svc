package getDocumentType

import "context"

type Repository interface {
	GetDocumentTypeRepo(ctx context.Context) ([]GetDocumentTypeResponse, error)
}

type Service interface {
	GetDocumentTypeSvc(ctx context.Context) ([]GetDocumentTypeResponse, error)
}

type GetDocumentTypeResponse struct {
	Id               int    `json:"id"`
	NameTypeDocument string `json:"nameTypeDocument"`
}
