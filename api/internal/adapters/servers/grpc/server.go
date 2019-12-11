package grpc

import (
	"calendar/api/internal/domain/entities"
	"calendar/api/internal/domain/interfaces"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
)


func StartServer(storage interfaces.EventStorage) {

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}


	server := grpc.NewServer()

	srv := &EventServiceServer{storage}

	entities.RegisterEventServiceServer(server, srv)

	// Start the server in a child routine
	go func() {
		if err := server.Serve(listener); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	fmt.Println("Server succesfully started on port 127.0.0.1:50051")
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	<-c
	fmt.Println("\nStopping the server...")
	server.Stop()
	listener.Close()
}


