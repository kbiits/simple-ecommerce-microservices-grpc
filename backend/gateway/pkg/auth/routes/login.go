package routes

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/kbiits/microservices-grpc-simple-ecommerce-gateway/pkg/auth/pb"
	"github.com/kbiits/microservices-grpc-simple-ecommerce-gateway/pkg/commons"
	"github.com/kbiits/microservices-grpc-simple-ecommerce-gateway/pkg/utils"
)

type LoginRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(ctx *gin.Context, client pb.AuthServiceClient) {
	var body LoginRequestBody

	if err := ctx.BindJSON(&body); err != nil {
		res := commons.NewBadRequestResponse(err)
		ctx.AbortWithStatusJSON(res.Code, res)
		return
	}

	res, err := client.Login(context.Background(), &pb.LoginRequest{
		Email:    body.Email,
		Password: body.Password,
	})

	if err != nil {
		resp := utils.ErrorResponseFromGrpc(err)
		ctx.AbortWithStatusJSON(resp.Code, resp)
		return
	} else if res.Response.Error != "" {
		ctx.JSON(int(res.Response.Status), utils.ErrorResponseFromMessage(res.Response.Error, res.Response.Status))
		return
	}

	ctx.JSON(int(res.Response.Status), utils.ToOkResponse(res.Response.Status, &gin.H{
		"token": res.Token,
	}))
}
