package routes

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/kbiits/microservices-grpc-simple-ecommerce-gateway/pkg/commons"
	"github.com/kbiits/microservices-grpc-simple-ecommerce-gateway/pkg/product/models"
	"github.com/kbiits/microservices-grpc-simple-ecommerce-gateway/pkg/product/pb"
	"github.com/kbiits/microservices-grpc-simple-ecommerce-gateway/pkg/utils"
)

type CreateProductRequestBody struct {
	Name  string `json:"name"`
	Sku   string `json:"sku"`
	Stock uint64 `json:"stock"`
	Price uint64 `json:"price"`
}

type CreateProductResponse struct {
	ProductId uint64 `json:"product_id"`
	CreateProductRequestBody
}

func CreateProduct(ctx *gin.Context, c pb.ProductServiceClient) {
	var body CreateProductRequestBody
	if err := ctx.BindJSON(&body); err != nil {
		res := commons.NewBadRequestResponse(err)
		ctx.AbortWithStatusJSON(res.Code, res)
		return
	}

	res, err := c.CreateProduct(context.Background(), &pb.CreateProductRequest{
		Product: &pb.Product{
			Name:  body.Name,
			Sku:   body.Sku,
			Stock: body.Stock,
			Price: body.Price,
		},
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
