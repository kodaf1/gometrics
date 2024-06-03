package main

import (
	"github.com/kodaf1/gometrics/internal/server"
	"github.com/kodaf1/gometrics/internal/storage"
)

func main() {
	srv := server.NewServer(storage.NewMapStorage())
	if err := srv.Run(); err != nil {
		panic(err)
	}
}
