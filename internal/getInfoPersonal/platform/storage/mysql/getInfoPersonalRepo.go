package mysql

import (
	"context"
	"database/sql"
	"errors"
	kitlog "github.com/go-kit/log"
	goconfig "github.com/iglin/go-config"
	"sgp-access-info-svc/internal/getInfoPersonal"
	"sgp-access-info-svc/kit/constants"
)

type GetInfoPersonalRepository struct {
	db     *sql.DB
	logger kitlog.Logger
}

func NewGetInfoPersonalRepository(db *sql.DB, logger kitlog.Logger) *GetInfoPersonalRepository {
	return &GetInfoPersonalRepository{db: db, logger: logger}
}

func (g *GetInfoPersonalRepository) GetInfoPersonalRepo(ctx context.Context, DocumentNumber string) (getInfoPersonal.GetInfoPersonalResponse, error) {

	config := goconfig.NewConfig("./application.yaml", goconfig.Yaml)

	statusActive := config.GetInt("app-properties.getComorbidity.idStatusActive")

	rowsSql := g.db.QueryRowContext(ctx, "select psipf.psi_first_name, psipf.psi_second_name, psipf.psi_first_last_name,\npsipf.psi_second_last_name, psipf.psi_date_of_birth, psipf.psi_document_number, psipf.psi_user, pssps.pss_gender_name, stdsd.std_state_name from \npsi_personal_information as psipf\ninner join pss_personal_sex as pssps\non pssps.pss_id = psipf.psi_sex_id\ninner join std_state_data as stdsd\non stdsd.std_id = psipf.psi_state_data_id\nwhere psi_document_number = ? AND psi_state_data_id = ?;", DocumentNumber, statusActive)
	var respDB SqlGetInfoPersonal
	if err := rowsSql.Scan(&respDB.FirstName, &respDB.SecondName, &respDB.LastName, &respDB.MothersLast, &respDB.DateOfBirth, &respDB.DocumentNumber, &respDB.User, &respDB.Sex, &respDB.StateData); err != nil {
		g.logger.Log("Data not found", constants.UUID, ctx.Value(constants.UUID))
	}

	resp := getInfoPersonal.GetInfoPersonalResponse{
		FirstName:       respDB.FirstName,
		SecondName:      respDB.SecondName,
		LastName:        respDB.LastName,
		MothersLastName: respDB.MothersLast,
		Sex:             respDB.Sex,
		DateOfBirth:     respDB.DateOfBirth,
		DocumentNumber:  respDB.DocumentNumber,
		User:            respDB.User,
	}

	if rowsSql != nil {
		g.logger.Log("Error -  Information could not be obtained", rowsSql.Err(), constants.UUID, ctx.Value(constants.UUID))
		return getInfoPersonal.GetInfoPersonalResponse{}, errors.New("Error: No data obteined")
	}

	g.logger.Log("Result Obtained", resp, constants.UUID, ctx.Value(constants.UUID))
	return resp, nil
}
