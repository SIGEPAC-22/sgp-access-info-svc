package mysql

type SqlGetTypeUser struct {
	Id       int64  `db:"tur_id"`
	NameType string `db:"tur_name_type"`
}
