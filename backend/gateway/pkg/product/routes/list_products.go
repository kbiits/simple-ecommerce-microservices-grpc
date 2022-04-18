package routes

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/kbiits/microservices-grpc-simple-ecommerce-gateway/pkg/product/models"
	"github.com/kbiits/microservices-grpc-simple-ecommerce-gateway/pkg/product/pb"
	"github.com/kbiits/microservices-grpc-simple-ecommerce-gateway/pkg/utils"
)

type Product struct {
}

func ListProducts(ctx *gin.Context, c pb.ProductServiceClient) {
	res, err := c.ListProducts(context.Background(), &pb.ListProductsRequest{})
	if err != nil {
		res := utils.ErrorResponseFromGrpc(err)
		ctx.AbortWithStatusJSON(res.Code, res)
		return
	} else if res.Response.Error != "" {
		ctx.JSON(int(res.Response.Status), utils.ErrorResponseFromMessage(res.Response.Error, res.Response.Status))
		return
	}

	var products []*models.Product
	for _, p := range res.Products {
		products = append(products, &models.Product{
			ProductId: p.ProductId,
			Name:      p.Name,
			Sku:       p.Sku,
			Stock:     p.Stock,
			Price:     p.Price,
		})
	}

	ctx.JSON(int(res.Response.Status), utils.ToOkResponse(res.Response.Status, products))
}
