package api

import (
	"context"
	"log"

	"github.com/arthurshafikov/events-collector/storage/internal/transport/grpc/generated"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewGRPCClient(ctx context.Context, address string) generated.CollectorClient {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		<-ctx.Done()
		conn.Close()
	}()

	return generated.NewCollectorClient(conn)
}
