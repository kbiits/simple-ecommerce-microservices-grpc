package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/kbiits/microservices-grpc-simple-ecommerce-gateway/pkg/auth/routes"
	"github.com/kbiits/microservices-grpc-simple-ecommerce-gateway/pkg/config"
)

func RegisterRoutes(router *gin.Engine, cfg *config.Config) (*AuthServiceClient, error) {
	svc, err := NewAuthServiceClient(cfg)
	if err != nil {
		return nil, err
	}

	authRoutes := router.Group("/auth")
	authRoutes.POST("/register", svc.Register)
	authRoutes.POST("/login", svc.Login)

	return svc, nil
}

func (svc *AuthServiceClient) Register(ctx *gin.Context) {
	routes.Register(ctx, svc.Client)
}

func (svc *AuthServiceClient) Login(ctx *gin.Context) {
	routes.Login(ctx, svc.Client)
}
