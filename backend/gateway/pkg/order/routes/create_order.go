package routes

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/kbiits/microservices-grpc-simple-ecommerce-gateway/pkg/commons"
	"github.com/kbiits/microservices-grpc-simple-ecommerce-gateway/pkg/order/models"
	"github.com/kbiits/microservices-grpc-simple-ecommerce-gateway/pkg/order/pb"
	"github.com/kbiits/microservices-grpc-simple-ecommerce-gateway/pkg/utils"
)

type CreateOrderRequestBody struct {
	ProductId uint64 `json:"product_id"`
	Quantity  uint64 `json:"quantity"`
	UserId    uint64 `json:"user_id"`
}

func CreateOrder(ctx *gin.Context, c pb.OrderServiceClient) {
	var body CreateOrderRequestBody

	if err := ctx.BindJSON(&body); err != nil {
		res := commons.NewBadRequestResponse(err)
		ctx.AbortWithStatusJSON(res.Code, res)
		return
	}

	res, err := c.CreateOrder(context.Background(), &pb.CreateOrderRequest{
		Order: &pb.Order{
			ProductId: body.ProductId,
			Quantity:  body.Quantity,
			UserId:    body.UserId,
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

	ctx.JSON(int(res.Response.Status), utils.ToOkResponse(res.Response.Status, &models.Order{
		ProductId: res.Order.ProductId,
		Quantity:  res.Order.Quantity,
		UserId:    res.Order.UserId,
		OrderId:   res.Order.OrderId,
	}))
}
