package order

import (
	"github.com/kbiits/microservices-grpc-simple-ecommerce-gateway/pkg/config"
	"github.com/kbiits/microservices-grpc-simple-ecommerce-gateway/pkg/order/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type OrderServiceClient struct {
	Client pb.OrderServiceClient
}

func NewOrderServiceClient(cfg *config.Config) (*OrderServiceClient, error) {
	cc, err := grpc.Dial(cfg.OrderServiceUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := pb.NewOrderServiceClient(cc)
	return &OrderServiceClient{
		Client: client,
	}, nil
}
