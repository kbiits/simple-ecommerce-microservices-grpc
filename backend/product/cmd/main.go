package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/kbiits/microservices-grpc-simple-ecommerce-product-service/pkg/config"
	"github.com/kbiits/microservices-grpc-simple-ecommerce-product-service/pkg/db"
	"github.com/kbiits/microservices-grpc-simple-ecommerce-product-service/pkg/pb"
	"github.com/kbiits/microservices-grpc-simple-ecommerce-product-service/pkg/services"
	"google.golang.org/grpc"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config, err : %v", err)
	}

	db, err := db.InitDB(cfg.PostgreDsn)
	if err != nil {
		log.Fatalf("Failed to connect to db, err : %v", err)
	}

	addr := cfg.Host + ":" + cfg.Port
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on addr %v, err : %v", addr, err)
	}

	srv := &services.Service{
		DB: db,
	}
	grpcServer := grpc.NewServer()
	pb.RegisterProductServiceServer(grpcServer, srv)

	log.Printf("Started grpc server on addr %v\n", addr)

	exitChan := make(chan int)
	go func() {
		signalChan := make(chan os.Signal, 1)
		signal.Notify(signalChan, os.Interrupt, syscall.SIGINT)
		signal := <-signalChan
		switch signal {
		case os.Interrupt:
			fallthrough
		case syscall.SIGINT:
			fallthrough
		case syscall.SIGTERM:
			exitChan <- 0
		default:
			exitChan <- 1
		}
	}()

	go func() {
		if grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to start grpc server, err : %v", err)
		}
	}()

	exitCode := <-exitChan
	log.Println("Server gracefully shutdown")
	grpcServer.GracefulStop()
	os.Exit(exitCode)
}
