package macinfo

import "github.com/jmoiron/sqlx"

type IRepository interface {
}

type repository struct {
	db *sqlx.DB
}

func NewMacInfoRepository(db *sqlx.DB) IRepository {
	return repository{
		db: db,
	}
}
