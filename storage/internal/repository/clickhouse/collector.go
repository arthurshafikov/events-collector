package clickhouse

import (
	"context"
	"fmt"

	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/arthurshafikov/events-collector/internal/core"
)

const EventsTable = "events"

type Collector struct {
	conn driver.Conn
}

func NewCollector(conn driver.Conn) *Collector {
	return &Collector{
		conn: conn,
	}
}

func (v *Collector) StoreEvents(ctx context.Context, events []core.Event) error {
	batch, err := v.conn.PrepareBatch(
		ctx,
		fmt.Sprintf("INSERT INTO %s (EventType, Time, UserIP)", EventsTable),
	)
	if err != nil {
		return err
	}

	for i := range events {
		if err := batch.AppendStruct(&events[i]); err != nil {
			return err
		}
	}

	return batch.Send()
}
