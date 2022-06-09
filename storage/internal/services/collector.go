package services

import (
	"context"

	"github.com/arthurshafikov/events-collector/internal/core"
	"github.com/arthurshafikov/events-collector/internal/repository"
)

type Collector struct {
	logger Logger
	repo   repository.Collector

	bufferSizeLimit int

	eventsBuffer []core.Event
}

func NewCollector(
	logger Logger,
	repo repository.Collector,
	bufferSizeLimit int,
) *Collector {
	return &Collector{
		logger: logger,
		repo:   repo,

		bufferSizeLimit: bufferSizeLimit,
	}
}

func (l *Collector) StoreEvent(ctx context.Context, event core.Event) error {
	l.eventsBuffer = append(l.eventsBuffer, event)

	if len(l.eventsBuffer) >= l.bufferSizeLimit {
		return l.FlushEvents(ctx)
	}

	return nil
}

func (l *Collector) FlushEvents(ctx context.Context) error {
	if err := l.repo.StoreEvents(ctx, l.eventsBuffer); err != nil {
		l.logger.Error(err)

		return core.ErrServerError
	}
	l.eventsBuffer = []core.Event{}

	return nil
}
