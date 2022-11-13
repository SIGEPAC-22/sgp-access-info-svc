package mysql

type SqlGetInfoPersonal struct {
	FirstName      string `db:"psi_first_name"`
	SecondName     string `db:"psi_second_name"`
	LastName       string `db:"psi_last_name"`
	MothersLast    string `db:"psi_mothers_last_name"`
	DateOfBirth    string `db:"psi_dateof_birth"`
	DocumentNumber string `db:"psi_document_number"`
	User           string `db:"cdt_user"`
	Sex            string `db:"pss_gender_name"`
	StateData      string `db:"std_state_name"`
}
