package mysql

import (
	"context"
	"database/sql"
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

func (g *GetInfoPersonalRepository) GetInfoPersonalRepo(ctx context.Context) ([]getInfoPersonal.GetInfoPersonalResponse, error) {

	config := goconfig.NewConfig("./application.yaml", goconfig.Yaml)

	rows, errDB := g.db.QueryContext(ctx, "SELECT psi_first_name, psi_second_name, psi_first_last_name, psi_second_last_name, pss_personal_sex.pss_gender_name, psi_date_of_birth, dct_document_name,\npsi_document_number, psi_user, tur_name_type, psi_account_creation_date, std_state_name from psi_personal_information\ninner join pss_personal_sex\non pss_personal_sex.pss_id = psi_personal_information.psi_sex_id\ninner join dct_document_type\non dct_document_type.dct_id = psi_personal_information.psi_document_type_id\ninner join tur_type_user\non tur_type_user.tur_id = psi_personal_information.psi_type_user\ninner join std_state_data\non std_state_data.std_id = psi_personal_information.psi_state_data_id;")
	if errDB != nil {
		g.logger.Log("Error while trying to get information for patient", constants.UUID, ctx.Value(constants.UUID))
		return []getInfoPersonal.GetInfoPersonalResponse{}, errDB
	}
	defer rows.Close()
	var resp []getInfoPersonal.GetInfoPersonalResponse
	for rows.Next() {
		var respDB SqlGetInfoPersonal
		if err := rows.Scan(&respDB.FirstName, &respDB.SecondName, &respDB.LastName, &respDB.SecondLastName, &respDB.Sex, &respDB.DateOfBirth, &respDB.TypeDocument, &respDB.DocumentNumber, &respDB.User, &respDB.TypeUser, &respDB.DateCreationAccount, &respDB.StateName); err != nil {
			g.logger.Log("error while trying to scan response from DB", "error", err.Error(), constants.UUID, ctx.Value(constants.UUID))
			return []getInfoPersonal.GetInfoPersonalResponse{}, err
		}
		resp = append(resp, getInfoPersonal.GetInfoPersonalResponse{
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
		})
	}

	if len(resp) == 0 {
		g.logger.Log("Data Not Found", constants.UUID, ctx.Value(constants.UUID))
	}
	return resp, nil
}
