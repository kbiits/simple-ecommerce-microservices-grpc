package routes

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kbiits/microservices-grpc-simple-ecommerce-gateway/pkg/auth/models"
	"github.com/kbiits/microservices-grpc-simple-ecommerce-gateway/pkg/auth/pb"
	"github.com/kbiits/microservices-grpc-simple-ecommerce-gateway/pkg/commons"
	"github.com/kbiits/microservices-grpc-simple-ecommerce-gateway/pkg/utils"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type RegisterRequestBody struct {
	FullName string    `json:"fullname"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Dob      time.Time `json:"dob"`
}

func Register(ctx *gin.Context, client pb.AuthServiceClient) {
	var body RegisterRequestBody

	if err := ctx.BindJSON(&body); err != nil {
		res := commons.NewBadRequestResponse(err)
		ctx.AbortWithStatusJSON(res.Code, res)
		return
	}

	res, err := client.Register(context.Background(), &pb.RegisterRequest{
		User: &pb.User{
			Email:    body.Email,
			Password: body.Password,
			Fullname: body.FullName,
			Dob:      timestamppb.New(body.Dob),
		},
	})
	if err != nil {
		resp := utils.ErrorResponseFromGrpc(err)
		ctx.AbortWithStatusJSON(resp.Code, resp)
		return
	} else if res.Response.Error != "" {
		ctx.JSON(int(res.Response.Status), utils.ErrorResponseFromMessage(res.Response.Error, res.Response.Status))
		return
	}

	ctx.JSON(int(res.Response.Status), utils.ToOkResponse(res.Response.Status, &models.User{
		UserId:   res.User.UserId,
		Email:    res.User.Email,
		Fullname: res.User.Fullname,
		Dob:      res.User.Dob.AsTime(),
	}))
}
