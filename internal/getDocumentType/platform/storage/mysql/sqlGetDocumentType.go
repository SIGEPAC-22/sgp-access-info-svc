package mysql

type SqlGetDocumentType struct {
	Id               int64  `db:"dct_id"`
	NameTypeDocument string `db:"dct_document_name"`
}
