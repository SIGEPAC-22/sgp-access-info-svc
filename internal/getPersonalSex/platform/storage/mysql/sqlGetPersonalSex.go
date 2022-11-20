package mysql

type SqlPersonalGetSex struct {
	Id      int64  `db:"pss_id"`
	NameSex string `db:"spt_gender_type"`
}
