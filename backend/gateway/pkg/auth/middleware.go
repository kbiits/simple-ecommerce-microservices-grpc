package auth

import (
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kbiits/microservices-grpc-simple-ecommerce-gateway/pkg/auth/pb"
	"github.com/kbiits/microservices-grpc-simple-ecommerce-gateway/pkg/utils"
)

type AuthMiddlewareConfig struct {
	service *AuthServiceClient
}

func NewAuthMiddleware(svc *AuthServiceClient) *AuthMiddlewareConfig {
	return &AuthMiddlewareConfig{
		service: svc,
	}
}

func (m *AuthMiddlewareConfig) RequireAuth(ctx *gin.Context) {
	bearer := ctx.GetHeader("authorization")

	if bearer == "" {
		res := utils.UnauthorizedResponse("Authorization header is empty")
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
		return
	}

	token := strings.TrimSpace(strings.Replace(bearer, "Bearer ", "", 1))
	if token == "" {
		res := utils.UnauthorizedResponse("Token not provided")
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
		return
	}

	res, err := m.service.Client.Authenticate(context.Background(), &pb.AuthenticateRequest{
		Token: token,
	})

	if err != nil || res.Response.Status != http.StatusOK {
		res := utils.UnauthorizedResponse("Invalid token")
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
		return
	} else if res.Response.Error != "" {
		ctx.JSON(int(res.Response.Status), utils.ErrorResponseFromMessage(res.Response.Error, res.Response.Status))
		return
	}

	ctx.Set("userId", res.UserId)
	ctx.Next()
}
