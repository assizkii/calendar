package grpc_server

import (
	"calendar/entities"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

func StartClient(requestCtx context.Context, requestOpts grpc.DialOption, port string) entities.EventServiceClient {

	// After Cobra root config init
	// We initialize the client
	fmt.Println("Starting Event Service Client")

	// Dial the server, returns a client connection
	conn, err := grpc.DialContext(requestCtx, port, requestOpts)

	if err != nil {
		log.Fatalf("Unable to establish client connection to : %s, %v", port, err)
	}

	// Instantiate the EventServiceClient with our client connection to the server
	return entities.NewEventServiceClient(conn)

}
