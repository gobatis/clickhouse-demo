package store

import (
	"database/sql"
	"github.com/gobatis/gobatis"
)

var (
	Migration = &migrateMapper{}
	Mapper    = &mapper{
		&logMapper{},
	}
)

func BindMapper(engine *gobatis.Engine) error {
	return engine.BindMapper(
		Migration,
		Mapper.logMapper,
	)
}

type mapper struct {
	*logMapper
}

type migrateMapper struct {
	Migrate func(db *gobatis.DB) error
}

type logMapper struct {
	InsertLog func(tx *sql.Tx, item *Log) error
	QueryLog  func() ([]*Log, error)
}
