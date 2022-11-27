package mysql

import (
	"context"
	"database/sql"
	kitlog "github.com/go-kit/log"
	goconfig "github.com/iglin/go-config"
	"sgp-access-info-svc/internal/getDocumentType"
	"sgp-access-info-svc/kit/constants"
)

type GetDocumentTypeRepo struct {
	db     *sql.DB
	logger kitlog.Logger
}

func NewGetDocumentTypeRepo(db *sql.DB, logger kitlog.Logger) *GetDocumentTypeRepo {
	return &GetDocumentTypeRepo{db: db, logger: logger}
}

func (g *GetDocumentTypeRepo) GetDocumentTypeRepo(ctx context.Context) ([]getDocumentType.GetDocumentTypeResponse, error) {

	config := goconfig.NewConfig("./application.yaml", goconfig.Yaml)
	id := config.GetInt("app-properties.getComorbidity.idStatusActive")
	rows, errDB := g.db.QueryContext(ctx, "SELECT dct_id, dct_document_name FROM dct_document_type WHERE dct_state_data_id = ?;", id)
	if errDB != nil {
		g.logger.Log("Error while trying to get document type for personal", constants.UUID, ctx.Value(constants.UUID))
		return []getDocumentType.GetDocumentTypeResponse{}, errDB
	}
	defer rows.Close()
	var resp []getDocumentType.GetDocumentTypeResponse
	for rows.Next() {
		var respDB SqlGetDocumentType
		if err := rows.Scan(&respDB.Id, &respDB.NameTypeDocument); err != nil {
			g.logger.Log("error while trying to scan response from DB", "error", err.Error(), constants.UUID, ctx.Value(constants.UUID))
			return []getDocumentType.GetDocumentTypeResponse{}, err
		}
		resp = append(resp, getDocumentType.GetDocumentTypeResponse{
			Id:               int(respDB.Id),
			NameTypeDocument: respDB.NameTypeDocument,
		})
	}

	if len(resp) == 0 {
		g.logger.Log("Data Not Found", constants.UUID, ctx.Value(constants.UUID))
	}
	return resp, nil
}
