package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/kbiits/microservices-grpc-simple-ecommerce-order-service/pkg/client"
	"github.com/kbiits/microservices-grpc-simple-ecommerce-order-service/pkg/config"
	"github.com/kbiits/microservices-grpc-simple-ecommerce-order-service/pkg/db"
	"github.com/kbiits/microservices-grpc-simple-ecommerce-order-service/pkg/pb/pb_order"
	"github.com/kbiits/microservices-grpc-simple-ecommerce-order-service/pkg/services"
	"google.golang.org/grpc"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config, err : %v", err)
	}

	db, err := db.InitDB(cfg)
	if err != nil {
		log.Fatalf("Failed to init db, err : %v", err)
	}
	log.Println("Connected to database")

	addr := cfg.Host + ":" + cfg.Port
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on %v, err : %v", addr, err)
	}

	productClient, err := client.NewProductServiceClient(cfg)
	if err != nil {
		log.Fatalf("Failed to init product client, err : %v", err)
	}
	srv := services.Service{
		DB:         db,
		ProductSvc: productClient,
	}
	grpcServer := grpc.NewServer()
	pb_order.RegisterOrderServiceServer(grpcServer, &srv)

	exitChan := make(chan int)

	go func() {
		if grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve grpc, err : %v", err)
		}
	}()

	log.Printf("GRPC server started on %v\n", addr)

	go func() {
		signalChan := make(chan os.Signal, 1)
		signal.Notify(signalChan, os.Interrupt, syscall.SIGINT)

		signal := <-signalChan
		switch signal {
		case os.Interrupt:
			fallthrough
		case syscall.SIGINT:
			exitChan <- 0
		default:
			exitChan <- 1
		}
	}()

	exitCode := <-exitChan
	log.Println("Gracefully shutdown server")
	grpcServer.GracefulStop()
	os.Exit(exitCode)
}
