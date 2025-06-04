package grpc

import (
	"log"

	"github.com/sunnysingha911/admin-service/gen/user-service/userpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials/insecure"
)

type UserClient struct {
	Client userpb.UserServiceClient
	Conn   *grpc.ClientConn
}

func NewUserClient(address string) *UserClient {
	clientConn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to create grpc client connection: %v", err)
	}

	// Optionally check connection state
	state := clientConn.GetState()
	if state != connectivity.Ready {
		log.Printf("Warning: gRPC connection state is %v", state)
	}

	return &UserClient{
		Client: userpb.NewUserServiceClient(clientConn),
		Conn:   clientConn,
	}
}
