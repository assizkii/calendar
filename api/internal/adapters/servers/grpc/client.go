package grpc

import (
	"calendar/api/internal/domain/entities"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

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