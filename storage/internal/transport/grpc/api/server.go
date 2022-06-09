package api

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/arthurshafikov/events-collector/internal/services"
	"github.com/arthurshafikov/events-collector/internal/transport/grpc/generated"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func RunGrpcServer(ctx context.Context, port string, logger services.Logger, services *services.Services) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalln(err)
	}

	collectorHandler := NewCollectorHandler(logger, services.Collectors)
	grpcServer := grpc.NewServer()
	generated.RegisterCollectorServer(grpcServer, collectorHandler)
	reflection.Register(grpcServer)

	go func() {
		<-ctx.Done()
		grpcServer.Stop()
	}()

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln(err)
	}
}
