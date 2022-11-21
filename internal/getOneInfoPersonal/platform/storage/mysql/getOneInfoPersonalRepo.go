package mysql

import (
	"context"
	"database/sql"
	"errors"
	kitlog "github.com/go-kit/log"
	goconfig "github.com/iglin/go-config"
	"sgp-access-info-svc/internal/getOneInfoPersonal"
	"sgp-access-info-svc/kit/constants"
)

type GetOneInfoPersonalRepository struct {
	db     *sql.DB
	logger kitlog.Logger
}

func NewGetOneInfoPersonalRepository(db *sql.DB, logger kitlog.Logger) *GetOneInfoPersonalRepository {
	return &GetOneInfoPersonalRepository{db: db, logger: logger}
}

func (g *GetOneInfoPersonalRepository) GetOneInfoPersonalRepo(ctx context.Context, id int) (getOneInfoPersonal.GetOneInfoPersonalResponse, error) {

	config := goconfig.NewConfig("./application.yaml", goconfig.Yaml)
	rowsSql := g.db.QueryRowContext(ctx, "SELECT psi_first_name, psi_second_name, psi_first_last_name, psi_second_last_name, pss_personal_sex.pss_gender_name, psi_date_of_birth, dct_document_name,\npsi_document_number, psi_user, tur_name_type, psi_account_creation_date, std_state_name from psi_personal_information\ninner join pss_personal_sex\non pss_personal_sex.pss_id = psi_personal_information.psi_sex_id\ninner join dct_document_type\non dct_document_type.dct_id = psi_personal_information.psi_document_type_id\ninner join tur_type_user\non tur_type_user.tur_id = psi_personal_information.psi_type_user\ninner join std_state_data\non std_state_data.std_id = psi_personal_information.psi_state_data_id\nwhere psi_id = ?;", id)
	var respDB SqlGetOneInfoPersonal
	if err := rowsSql.Scan(&respDB.FirstName, &respDB.SecondName, &respDB.LastName, &respDB.SecondLastName, &respDB.Sex, &respDB.DateOfBirth, &respDB.TypeDocument, &respDB.DocumentNumber, &respDB.User, &respDB.TypeUser, &respDB.DateCreationAccount, &respDB.StateName); err != nil {
		g.logger.Log("Data not found", constants.UUID, ctx.Value(constants.UUID))
	}

	resp := getOneInfoPersonal.GetOneInfoPersonalResponse{
		FirstName:           respDB.FirstName,
		SecondName:          respDB.SecondName,
		LastName:            respDB.LastName,
		SecondLastName:      respDB.SecondLastName,
		Sex:                 respDB.Sex,
		DateOfBirth:         respDB.DateOfBirth.Format(config.GetString("app-properties.getInfoPatient.dateBirth-Format")),
		TypeDocument:        respDB.TypeDocument,
		DocumentNumber:      respDB.DocumentNumber,
		User:                respDB.User,
		TypeUser:            respDB.TypeUser,
		DateCreationAccount: respDB.DateCreationAccount.Format(config.GetString("app-properties.getInfoPatient.dateBirth-Format")),
		StateName:           respDB.StateName,
	}

	if rowsSql == nil {
		g.logger.Log("Error -  Information could not be obtained", rowsSql.Err(), constants.UUID, ctx.Value(constants.UUID))
		return getOneInfoPersonal.GetOneInfoPersonalResponse{}, errors.New("Error: No data obteined")
	}

	g.logger.Log("Result Obtained", resp, constants.UUID, ctx.Value(constants.UUID))
	return resp, nil
}
