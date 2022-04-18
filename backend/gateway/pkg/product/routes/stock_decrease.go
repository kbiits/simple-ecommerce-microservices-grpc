package routes

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/kbiits/microservices-grpc-simple-ecommerce-gateway/pkg/commons"
	"github.com/kbiits/microservices-grpc-simple-ecommerce-gateway/pkg/product/pb"
	"github.com/kbiits/microservices-grpc-simple-ecommerce-gateway/pkg/utils"
)

type SetStockRequestBody struct {
	ProductId uint64 `json:"product_id"`
	OrderId   uint64 `json:"order_id"`
}

func StockDecrease(ctx *gin.Context, c pb.ProductServiceClient) {
	var body SetStockRequestBody
	if err := ctx.BindJSON(&body); err != nil {
		res := commons.NewBadRequestResponse(err)
		ctx.JSON(res.Code, &res)
		return
	}

	res, err := c.DecreaseStock(context.Background(), &pb.DecreaseStockRequest{
		ProductId: body.ProductId,
		OrderId: body.OrderId,
	})

	if err != nil {
		res := utils.ErrorResponseFromGrpc(err)
		ctx.JSON(res.Code, res)
		return
	} else if res.Response.Error != "" {
		ctx.JSON(int(res.Response.Status), utils.ErrorResponseFromMessage(res.Response.Error, res.Response.Status))
		return
	}

	ctx.JSON(int(res.Response.Status), utils.ToOkResponse(res.Response.Status, nil))
}
