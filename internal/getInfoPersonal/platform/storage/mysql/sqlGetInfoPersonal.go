package mysql

import "time"

type SqlGetInfoPersonal struct {
	Id                  int       `db:"psi_id"`
	FirstName           string    `db:"psi_first_name"`
	SecondName          string    `db:"psi_second_name"`
	LastName            string    `db:"psi_first_last_name"`
	SecondLastName      string    `db:"psi_second_last_name"`
	Sex                 string    `db:"pss_gender_name"`
	DateOfBirth         time.Time `db:"psi_date_of_birth"`
	TypeDocument        string    `db:"dct_document_name"`
	DocumentNumber      string    `db:"psi_document_number"`
	User                string    `db:"psi_user"`
	TypeUser            string    `db:"tur_name_type"`
	DateCreationAccount time.Time `db:"psi_account_creation_date"`
	StateName           string    `db:"std_state_name"`
}
