package utils

import (
	"github.com/kbiits/microservices-grpc-simple-ecommerce-auth-service/pkg/models"
	"github.com/kbiits/microservices-grpc-simple-ecommerce-auth-service/pkg/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func UserPbToUserModels(user *pb.User) *models.User {
	return &models.User{
		UserId:   user.UserId,
		Email:    user.Email,
		Password: user.Password,
	}
}

func UserModelsToUserPb(user *models.User) *pb.User {
	return &pb.User{
		UserId:   user.UserId,
		Email:    user.Email,
		Password: user.Password,
		Fullname: user.Fullname,
		Dob:      timestamppb.New(user.Dob),
	}
}
