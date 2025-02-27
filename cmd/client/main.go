package main

import (
	"context"

	"github.com/fsaintjacques/echo/client"
)

func main() {
	client.Client(context.Background(), "localhost:8080")
}
