package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

func main()  {
	listener, err := net.Listen("tcp", "8080")
	if err != nil {
		log.Fatalf("unable to listen on 8080 port: %v", err)
	}

	srv := grpc.NewServer()
}