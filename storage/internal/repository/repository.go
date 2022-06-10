package repository

import (
	"context"

	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/arthurshafikov/events-collector/storage/internal/core"
	"github.com/arthurshafikov/events-collector/storage/internal/repository/clickhouse"
)

type Collector interface {
	StoreEvents(ctx context.Context, events []core.Event) error
}

type Repository struct {
	Collector
}

func NewRepository(conn driver.Conn) *Repository {
	return &Repository{
		Collector: clickhouse.NewCollector(conn),
	}
}
