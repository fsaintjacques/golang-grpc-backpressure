package client

import (
	"context"
	"log"
	"time"

	"github.com/fsaintjacques/echo/proto/gen/echoservice"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Client(ctx context.Context, host string) {
	conn, err := grpc.Dial(host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := echoservice.NewEchoServiceClient(conn)
	res, err := client.Echo(ctx, &echoservice.Request{})
	if err != nil {
		panic(err)
	}

	n := time.Duration(1)

	for {
		time.Sleep(n * time.Second)

		log.Printf("Receiving message")
		_, err := res.Recv()
		if err != nil {
			panic(err)
		}
	}
}
