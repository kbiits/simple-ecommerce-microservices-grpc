package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kbiits/microservices-grpc-simple-ecommerce-gateway/pkg/auth"
	"github.com/kbiits/microservices-grpc-simple-ecommerce-gateway/pkg/config"
	"github.com/kbiits/microservices-grpc-simple-ecommerce-gateway/pkg/order"
	"github.com/kbiits/microservices-grpc-simple-ecommerce-gateway/pkg/product"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config, err : %v", err)
	}

	router := gin.Default()
	router.Use(cors.Default())

	authSvc, err := auth.RegisterRoutes(router, cfg)
	if err != nil {
		log.Fatalf("Failed to create auth service, err : %v", err)
	}

	authMiddleware := auth.NewAuthMiddleware(authSvc)

	_, err = order.RegisterRoutes(router, cfg, authMiddleware)
	if err != nil {
		log.Fatalf("Failed to create order service, err : %v", err)
	}

	_, err = product.RegisterRoutes(router, cfg, authMiddleware)
	if err != nil {
		log.Fatalf("Failed to create product service, err : %v", err)
	}

	addr := cfg.Host + ":" + cfg.Port

	exitChan := make(chan int, 1)

	srv := http.Server{
		Addr:    addr,
		Handler: router,
	}

	go func() {
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("failed to start server, err : %v", err)
		}
	}()

	go func(exitChan chan int) {
		signalChan := make(chan os.Signal, 1)
		signal.Notify(signalChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		signal := <-signalChan
		close(signalChan)
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
	}(exitChan)

	exitCode := <-exitChan
	err = srv.Shutdown(context.Background())
	if err != nil {
		log.Fatalf("failed to gracefully shutdown the server, err : %v", err)
	}

	fmt.Println("Server grafecully shutdown")
	fmt.Println("exit code :", exitCode)
	close(exitChan)
	os.Exit(exitCode)
}
