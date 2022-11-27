package mysql

import (
	"context"
	"database/sql"
	kitlog "github.com/go-kit/log"
	goconfig "github.com/iglin/go-config"
	"sgp-access-info-svc/internal/getTypeUser"
	"sgp-access-info-svc/kit/constants"
)

type GetTypeUserRepo struct {
	db     *sql.DB
	logger kitlog.Logger
}

func NewGetTypeUserRepo(db *sql.DB, logger kitlog.Logger) *GetTypeUserRepo {
	return &GetTypeUserRepo{db: db, logger: logger}
}

func (g *GetTypeUserRepo) GetTypeUserRepo(ctx context.Context) ([]getTypeUser.GetTypeUserResponse, error) {
	config := goconfig.NewConfig("./application.yaml", goconfig.Yaml)
	id := config.GetInt("app-properties.getComorbidity.idStatusActive")
	rows, errDB := g.db.QueryContext(ctx, "SELECT tur_id, tur_name_type FROM tur_type_user WHERE tur_state_data_id = ?;", id)
	if errDB != nil {
		g.logger.Log("Error while trying to get sex for personal", constants.UUID, ctx.Value(constants.UUID))
		return []getTypeUser.GetTypeUserResponse{}, errDB
	}
	defer rows.Close()
	var resp []getTypeUser.GetTypeUserResponse
	for rows.Next() {
		var respDB SqlGetTypeUser
		if err := rows.Scan(&respDB.Id, &respDB.NameType); err != nil {
			g.logger.Log("error while trying to scan response from DB", "error", err.Error(), constants.UUID, ctx.Value(constants.UUID))
			return []getTypeUser.GetTypeUserResponse{}, err
		}
		resp = append(resp, getTypeUser.GetTypeUserResponse{
			Id:       int(respDB.Id),
			NameType: respDB.NameType,
		})
	}
	if len(resp) == 0 {
		g.logger.Log("Data Not Found", constants.UUID, ctx.Value(constants.UUID))
	}
	return resp, nil
}
