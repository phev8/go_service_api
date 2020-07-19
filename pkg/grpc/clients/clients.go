package clients

import (
	"log"

	apiA "github.com/phev8/go_service_A/pkg/api"
	apiB "github.com/phev8/go_service_B/pkg/api"
	"google.golang.org/grpc"
)

func connectToGRPCServer(addr string) *grpc.ClientConn {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect to %s: %v", addr, err)
	}
	return conn
}

func ConnectToAService(addr string) (client apiA.ServiceAClient, close func() error) {
	serverConn := connectToGRPCServer(addr)
	return apiA.NewServiceAClient(serverConn), serverConn.Close
}

func ConnectToBService(addr string) (client apiB.ServiceBClient, close func() error) {
	serverConn := connectToGRPCServer(addr)
	return apiB.NewServiceBClient(serverConn), serverConn.Close
}
