package auth

import (
	"github.com/kbiits/microservices-grpc-simple-ecommerce-gateway/pkg/auth/pb"
	"github.com/kbiits/microservices-grpc-simple-ecommerce-gateway/pkg/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AuthServiceClient struct {
	Client pb.AuthServiceClient
}

func NewAuthServiceClient(cfg *config.Config) (*AuthServiceClient, error) {
	cc, err := grpc.Dial(cfg.AuthServiceUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := pb.NewAuthServiceClient(cc)
	return &AuthServiceClient{
		Client: client,
	}, nil
}
