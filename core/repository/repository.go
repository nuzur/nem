package repository

import (
	"context"

	"database/sql"

	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/core/repository/list"
)

type Implementation struct {
	Queries *nemdb.Queries
	DB      *sql.DB
	List    *list.Implementation
}

func New(db *sql.DB) *Implementation {
	queries := nemdb.New(db)
	return &Implementation{
		Queries: queries,
		DB:      db,
		List:    list.New(),
	}
}

func (i *Implementation) BuildListEntityQuery(ctx context.Context, request list.ListRequest, entity list.ListEntity, onlyCount bool, includeColumns []string, excludeColumns []string) (string, error) {
	return i.List.BuildListEntityQuery(ctx, request, entity, onlyCount, includeColumns, excludeColumns)
}
