package dataaccess

import (
	"database/sql"
	"student-service/pkg/data-access/dto"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)

func InitializeSequelDB(connString string) *bun.DB {
	// initialize pgx config
	config, err := pgx.ParseConfig(connString)
	if err != nil {
		panic(err)
	}

	// use simple protocol as bun does not benefit from using implicit prepared statements
	config.DefaultQueryExecMode = pgx.QueryExecModeSimpleProtocol

	var sqlDB *sql.DB = stdlib.OpenDB(*config)

	db := bun.NewDB(sqlDB, pgdialect.New())
	db.RegisterModel((*dto.ClassToStudent)(nil))
	db.RegisterModel((*dto.ClassToTeacher)(nil))

	return db
}
