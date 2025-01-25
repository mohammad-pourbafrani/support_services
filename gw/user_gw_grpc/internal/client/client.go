package client


import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func GrpcClientServerConnection(server_address string) *grpc.ClientConn {
	conn, err := grpc.NewClient(server_address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	return conn
}
