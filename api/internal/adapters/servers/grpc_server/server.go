package grpc_server

import (
	"calendar/api/internal/adapters/storages/inmemory"
	"calendar/api/internal/domain/entities"
	"calendar/api/internal/utils"
	"fmt"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
)


func StartServer() {
	fmt.Println("Starting Calendar Service Server")
	appConf := utils.GetAppConfig()

	listener, err := net.Listen("tcp", appConf.Host)
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}

	//db := database.GetDBConnection()
	//database.RunMigrations(db)
	//defer db.Close()

	logger, err := utils.InitLogger(appConf)
	if err != nil {
		log.Fatalf("Failed to init logger: %v", err)
	}


	server := grpc.NewServer(grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
		grpc_ctxtags.UnaryServerInterceptor(),
		grpc_opentracing.UnaryServerInterceptor(),
		grpc_zap.UnaryServerInterceptor(logger),
		grpc_recovery.UnaryServerInterceptor(),
	)),)

	storage :=  inmemory.New()

	srv := &EventServiceServer{storage}

	entities.RegisterEventServiceServer(server, srv)

	// Start the server in a child routine
	go func() {
		if err := server.Serve(listener); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	fmt.Printf("Server succesfully started on port %s", appConf.Host)
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	<-c
	fmt.Println("\nStopping the server...")
	server.Stop()
	listener.Close()
}
