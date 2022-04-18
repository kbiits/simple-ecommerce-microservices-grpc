package order

import (
	"github.com/gin-gonic/gin"
	"github.com/kbiits/microservices-grpc-simple-ecommerce-gateway/pkg/auth"
	"github.com/kbiits/microservices-grpc-simple-ecommerce-gateway/pkg/config"
	"github.com/kbiits/microservices-grpc-simple-ecommerce-gateway/pkg/order/routes"
)

func RegisterRoutes(router *gin.Engine, cfg *config.Config, authMid *auth.AuthMiddlewareConfig) (*OrderServiceClient, error) {
	svc, err := NewOrderServiceClient(cfg)
	if err != nil {
		return nil, err
	}

	orderRoutes := router.Group("/order")
	orderRoutes.Use(authMid.RequireAuth)

	orderRoutes.POST("", svc.createOrder)

	return svc, nil
}

func (svc *OrderServiceClient) createOrder(ctx *gin.Context) {
	routes.CreateOrder(ctx, svc.Client)
}
