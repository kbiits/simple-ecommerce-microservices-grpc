package client

import (
	"context"

	"github.com/kbiits/microservices-grpc-simple-ecommerce-order-service/pkg/config"
	"github.com/kbiits/microservices-grpc-simple-ecommerce-order-service/pkg/pb/pb_product"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ProductServiceClient struct {
	Client pb_product.ProductServiceClient
}

func NewProductServiceClient(cfg *config.Config) (*ProductServiceClient, error) {
	cc, err := grpc.Dial(cfg.ProductSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, err
	}

	return &ProductServiceClient{
		pb_product.NewProductServiceClient(cc),
	}, nil
}

func (c *ProductServiceClient) FindOneProduct(productId uint64) (*pb_product.FindOneResponse, error) {
	return c.Client.FindOne(context.Background(), &pb_product.FindOneRequest{
		ProductId: productId,
	})
}

func (c *ProductServiceClient) DecreaseProductStock(productId, orderId uint64) (*pb_product.DecreaseStockResponse, error) {
	return c.Client.DecreaseStock(context.Background(), &pb_product.DecreaseStockRequest{
		ProductId: productId,
		OrderId:   orderId,
	})
}
