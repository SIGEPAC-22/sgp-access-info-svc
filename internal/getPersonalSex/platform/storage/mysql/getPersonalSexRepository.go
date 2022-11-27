package mysql

import (
	"context"
	"database/sql"
	kitlog "github.com/go-kit/log"
	goconfig "github.com/iglin/go-config"
	"sgp-access-info-svc/internal/getPersonalSex"
	"sgp-access-info-svc/kit/constants"
)

type GetPersonalSexRepo struct {
	db     *sql.DB
	logger kitlog.Logger
}

func NewGetPersonalSexRepo(db *sql.DB, logger kitlog.Logger) *GetPersonalSexRepo {
	return &GetPersonalSexRepo{db: db, logger: logger}
}

func (g *GetPersonalSexRepo) GetPersonalSexRepo(ctx context.Context) ([]getPersonalSex.GetPersonalSexResponse, error) {
	config := goconfig.NewConfig("./application.yaml", goconfig.Yaml)
	id := config.GetInt("app-properties.getComorbidity.idStatusActive")
	rows, errDB := g.db.QueryContext(ctx, "SELECT pss_id, pss_gender_name FROM pss_personal_sex WHERE pss_state_data_id = ?;", id)
	if errDB != nil {
		g.logger.Log("Error while trying to get sex for personal", constants.UUID, ctx.Value(constants.UUID))
		return []getPersonalSex.GetPersonalSexResponse{}, errDB
	}
	defer rows.Close()
	var resp []getPersonalSex.GetPersonalSexResponse
	for rows.Next() {
		var respDB SqlPersonalGetSex
		if err := rows.Scan(&respDB.Id, &respDB.NameSex); err != nil {
			g.logger.Log("error while trying to scan response from DB", "error", err.Error(), constants.UUID, ctx.Value(constants.UUID))
			return []getPersonalSex.GetPersonalSexResponse{}, err
		}
		resp = append(resp, getPersonalSex.GetPersonalSexResponse{
			Id:      int(respDB.Id),
			NameSex: respDB.NameSex,
		})
	}
	if len(resp) == 0 {
		g.logger.Log("Data Not Found", constants.UUID, ctx.Value(constants.UUID))
	}
	return resp, nil
}
