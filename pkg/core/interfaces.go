package core

import (
	"context"

	"github.com/ezz-amine/Jadwal/pkg/sqlc"
)

type EntriesLister interface {
	ListEntries(ctx context.Context, tableID int64) ([]sqlc.TodoEntry, error)
}

type TablesLister interface {
	EntriesLister

	ListTables(ctx context.Context) ([]sqlc.TodoTable, error)
}
