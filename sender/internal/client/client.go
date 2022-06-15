package client

import (
	"context"
	"log"

	"github.com/arthurshafikov/events-collector/sender/internal/client/generated"
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
