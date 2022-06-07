package api

import (
	"context"
	"log"
	"net"

	"github.com/arthurshafikov/events-collector/internal/services"
	"github.com/arthurshafikov/events-collector/internal/transport/grpc/generated"
	"google.golang.org/grpc"
)

func RunGrpcServer(ctx context.Context, address string, logger services.Logger, services *services.Services) {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalln(err)
	}

	collectorHandler := NewCollectorHandler(logger, services.Collectors)
	grpcServer := grpc.NewServer()
	generated.RegisterCollectorServer(grpcServer, collectorHandler)

	go func() {
		<-ctx.Done()
		grpcServer.Stop()
	}()

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln(err)
	}
}
