package api

import (
	"context"
	"time"

	"github.com/arthurshafikov/events-collector/internal/core"
	"github.com/arthurshafikov/events-collector/internal/services"
	"github.com/arthurshafikov/events-collector/internal/transport/grpc/generated"
)

var successResponse = &generated.ServerResponse{
	Data: "OK",
}

type CollectorService struct {
	logger           services.Logger
	collectorService services.Collectors

	generated.UnimplementedCollectorServer
}

func NewCollectorHandler(logger services.Logger, collectorService services.Collectors) *CollectorService {
	return &CollectorService{
		logger:           logger,
		collectorService: collectorService,
	}
}

func (c *CollectorService) StoreEvent(
	ctx context.Context,
	req *generated.EventRequest,
) (*generated.ServerResponse, error) {
	if err := c.collectorService.StoreEvent(ctx, core.Event{
		EventType: req.EventType,
		Time:      time.Now(),
		UserIP:    req.UserIP,
	}); err != nil {
		c.logger.Error(err)

		return nil, core.ErrServerError
	}

	return successResponse, nil
}
