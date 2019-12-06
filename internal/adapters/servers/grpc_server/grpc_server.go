package grpc_server

import (
	"calendar/internal/domain/entities"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
)


func StartServer() {

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}


	server := grpc.NewServer()

	srv := &EventServiceServer{}

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

func StartClient(requestCtx context.Context, requestOpts grpc.DialOption) entities.EventServiceClient {

	// After Cobra root config init
	// We initialize the client
	fmt.Println("Starting Event Service Client")


	// Dial the server, returns a client connection
	conn, err := grpc.DialContext(requestCtx,"127.0.0.1:50051", requestOpts)

	if err != nil {
		log.Fatalf("Unable to establish client connection to 127.0.0.1:50052: %v", err)
	}

	// Instantiate the EventServiceClient with our client connection to the server
	return entities.NewEventServiceClient(conn)

}
