package server

import (
	"crypto/rand"
	"log"

	"github.com/fsaintjacques/echo/proto/gen/echoservice"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// server struct holds the gRPC server instance and implements the StockServiceServer interface.
type server struct {
	echoservice.UnimplementedEchoServiceServer
	GrpcSrv *grpc.Server
}

func New() *server {
	grpcServer := grpc.NewServer()
	srv := &server{
		GrpcSrv: grpcServer,
	}

	// Register the StockService with the gRPC server instance.
	echoservice.RegisterEchoServiceServer(grpcServer, srv)

	return srv
}

// GetUpdates streams stock updates to the client. It creates a stock with a starting price and sends
// random updates to the connected client every second.
func (s *server) Echo(_ *echoservice.Request, stream grpc.ServerStreamingServer[echoservice.Response]) error {
	payload := make([]byte, 4*1024*1024-32)
	if _, err := rand.Read(payload); err != nil {
		return status.Error(codes.Unknown, "failed to generate random payload: "+err.Error())
	}

	response := &echoservice.Response{
		Payload: payload,
	}

	for {
		log.Printf("Sending message")
		if err := stream.Send(response); err != nil {
			return status.Error(codes.Unknown, "failed to send update to client: "+err.Error())
		}
	}
}
