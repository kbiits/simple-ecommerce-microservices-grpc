package product

import (
	"github.com/kbiits/microservices-grpc-simple-ecommerce-gateway/pkg/config"
	"github.com/kbiits/microservices-grpc-simple-ecommerce-gateway/pkg/product/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ProductServiceClient struct {
	Client pb.ProductServiceClient
}

func NewProductServiceClient(cfg *config.Config) (*ProductServiceClient, error) {
	cc, err := grpc.Dial(cfg.ProductServiceUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := pb.NewProductServiceClient(cc)

	return &ProductServiceClient{
		Client: client,
	}, nil
}
