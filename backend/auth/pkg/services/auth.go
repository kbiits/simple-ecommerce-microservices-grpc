package services

import (
	"context"
	"net/http"

	"github.com/kbiits/microservices-grpc-simple-ecommerce-auth-service/pkg/db"
	"github.com/kbiits/microservices-grpc-simple-ecommerce-auth-service/pkg/models"
	"github.com/kbiits/microservices-grpc-simple-ecommerce-auth-service/pkg/pb"
	"github.com/kbiits/microservices-grpc-simple-ecommerce-auth-service/pkg/utils"
)

type Service struct {
	pb.UnimplementedAuthServiceServer
	DB  *db.DB
	Jwt utils.JwtWrapper
}

var _ pb.AuthServiceServer = &Service{}

func (s *Service) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	var user models.User

	if result := s.DB.Where(&models.User{Email: req.User.Email}).First(&user); result.Error == nil {
		return &pb.RegisterResponse{
			Response: &pb.Response{
				Status: http.StatusConflict,
				Error:  "Email already exists",
			},
		}, nil
	}

	user.Email = req.User.Email
	user.Fullname = req.User.Fullname
	user.Dob = req.User.Dob.AsTime()
	hashed, err := utils.HashPassword(req.User.Password)
	if err != nil {
		return &pb.RegisterResponse{
			Response: &pb.Response{
				Error:  "Failed to prepare user data",
				Status: http.StatusInternalServerError,
			},
		}, nil
	}
	user.Password = hashed

	s.DB.Create(&user)

	return &pb.RegisterResponse{
		Response: &pb.Response{
			Status: http.StatusCreated,
		},
		User: utils.UserModelsToUserPb(&user),
	}, nil
}

func (s *Service) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	var user models.User

	res := s.DB.Find(&models.User{Email: req.Email}).First(&user)
	if res.Error != nil {
		return &pb.LoginResponse{
			Response: &pb.Response{
				Status: http.StatusNotFound,
				Error:  "User not found",
			},
		}, nil
	}

	if !utils.ComparePasswordAndHash(user.Password, req.Password) {
		return &pb.LoginResponse{
			Response: &pb.Response{
				Status: http.StatusNotFound,
				Error:  "User not found",
			},
		}, nil
	}

	token, err := s.Jwt.GenerateToken(user)
	if err != nil {
		return &pb.LoginResponse{
			Response: &pb.Response{
				Error:  "Failed to prepare response",
				Status: http.StatusInternalServerError,
			},
		}, nil
	}

	return &pb.LoginResponse{
		Response: &pb.Response{
			Status: http.StatusOK,
		},
		Token: token,
	}, nil
}

func (s *Service) Authenticate(ctx context.Context, req *pb.AuthenticateRequest) (*pb.AuthenticateResponse, error) {
	claims, err := s.Jwt.VerifyToken(req.Token)
	if err != nil {
		return &pb.AuthenticateResponse{
			Response: &pb.Response{
				Error:  err.Error(),
				Status: http.StatusBadRequest,
			},
		}, nil
	}

	var user models.User
	if result := s.DB.Where(&models.User{UserId: claims.Id}).First(&user); result.Error != nil {
		return &pb.AuthenticateResponse{
			Response: &pb.Response{
				Error:  "User not found",
				Status: http.StatusNotFound,
			},
		}, nil
	}

	return &pb.AuthenticateResponse{
		Response: &pb.Response{
			Status: http.StatusOK,
		},
		UserId: user.UserId,
	}, nil
}
