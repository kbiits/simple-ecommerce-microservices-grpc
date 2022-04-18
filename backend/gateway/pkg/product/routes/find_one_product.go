package routes

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/kbiits/microservices-grpc-simple-ecommerce-gateway/pkg/commons"
	"github.com/kbiits/microservices-grpc-simple-ecommerce-gateway/pkg/product/models"
	"github.com/kbiits/microservices-grpc-simple-ecommerce-gateway/pkg/product/pb"
	"github.com/kbiits/microservices-grpc-simple-ecommerce-gateway/pkg/utils"
)

type FindOneProductRequest struct {
	ProductId uint64 `uri:"product_id"`
}

func FindOneProduct(ctx *gin.Context, c pb.ProductServiceClient) {
	var body FindOneProductRequest
	if err := ctx.BindUri(&body); err != nil {
		res := commons.NewBadRequestResponse(err)
		ctx.AbortWithStatusJSON(res.Code, res)
		return
	}
	
	res, err := c.FindOne(context.Background(), &pb.FindOneRequest{
		ProductId: body.ProductId,
	})

	if err != nil {
		res := utils.ErrorResponseFromGrpc(err)
		ctx.AbortWithStatusJSON(res.Code, res)
		return
	} else if res.Response.Error != "" {
		ctx.JSON(int(res.Response.Status), utils.ErrorResponseFromMessage(res.Response.Error, res.Response.Status))
		return
	}

	ctx.JSON(int(res.Response.Status), utils.ToOkResponse(res.Response.Status, &models.Product{
		ProductId: res.Product.ProductId,
		Name:      res.Product.Name,
		Sku:       res.Product.Sku,
		Stock:     res.Product.Stock,
		Price:     res.Product.Stock,
	}))
}
