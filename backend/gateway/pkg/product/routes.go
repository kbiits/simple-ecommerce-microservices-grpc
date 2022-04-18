package product

import (
	"github.com/gin-gonic/gin"
	"github.com/kbiits/microservices-grpc-simple-ecommerce-gateway/pkg/auth"
	"github.com/kbiits/microservices-grpc-simple-ecommerce-gateway/pkg/config"
	"github.com/kbiits/microservices-grpc-simple-ecommerce-gateway/pkg/product/routes"
)

func RegisterRoutes(router *gin.Engine, cfg *config.Config, authMid *auth.AuthMiddlewareConfig) (*ProductServiceClient, error) {
	svc, err := NewProductServiceClient(cfg)
	if err != nil {
		return nil, err
	}

	productRoutes := router.Group("/product")
	productRoutes.Use(authMid.RequireAuth)
	productRoutes.GET("/:product_id", svc.findOneProduct)
	productRoutes.POST("", svc.createProduct)

	return svc, nil
}

func (svc *ProductServiceClient) findOneProduct(ctx *gin.Context) {
	routes.FindOneProduct(ctx, svc.Client)
}

func (svc *ProductServiceClient) createProduct(ctx *gin.Context) {
	routes.CreateProduct(ctx, svc.Client)
}
