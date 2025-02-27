package main

import (
	"net"

	"github.com/fsaintjacques/echo/server"
)

func main() {
	server := server.New()
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	server.GrpcSrv.Serve(l)
}
