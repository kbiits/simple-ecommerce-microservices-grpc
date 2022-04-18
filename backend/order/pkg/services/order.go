package services

import (
	"context"
	"net/http"

	"github.com/kbiits/microservices-grpc-simple-ecommerce-order-service/pkg/client"
	"github.com/kbiits/microservices-grpc-simple-ecommerce-order-service/pkg/db"
	"github.com/kbiits/microservices-grpc-simple-ecommerce-order-service/pkg/models"
	"github.com/kbiits/microservices-grpc-simple-ecommerce-order-service/pkg/pb/pb_order"
)

type Service struct {
	pb_order.UnimplementedOrderServiceServer
	DB         *db.DB
	ProductSvc *client.ProductServiceClient
}

func (s *Service) CreateOrder(ctx context.Context, req *pb_order.CreateOrderRequest) (*pb_order.CreateOrderResponse, error) {
	product, err := s.ProductSvc.FindOneProduct(req.Order.ProductId)
	if err != nil {
		return &pb_order.CreateOrderResponse{
			Response: &pb_order.Response{
				Status: http.StatusBadRequest,
				Error:  err.Error(),
			},
		}, nil
	} else if product.Response.Status >= http.StatusBadRequest {
		return &pb_order.CreateOrderResponse{
			Response: &pb_order.Response{
				Status: product.Response.Status,
				Error:  product.Response.Error,
			},
		}, nil
	} else if product.Product.Stock < req.Order.Quantity {
		return &pb_order.CreateOrderResponse{
			Response: &pb_order.Response{
				Status: http.StatusConflict, Error: "Stock too low",
			},
		}, nil
	}

	order := models.Order{
		ProductId: product.Product.ProductId,
		Price:     product.Product.Price,
		UserId:    req.Order.UserId,
		Quantity:  req.Order.Quantity,
	}

	res := s.DB.Create(&order)
	if res.Error != nil {
		return &pb_order.CreateOrderResponse{
			Response: &pb_order.Response{
				Status: http.StatusInternalServerError, Error: "Failed to create order",
			},
		}, nil
	}

	decreaseResp, err := s.ProductSvc.DecreaseProductStock(order.ProductId, order.Id)
	if err != nil || decreaseResp.Response.Status >= http.StatusBadRequest {
		res.Delete(&order)
		return &pb_order.CreateOrderResponse{
			Response: &pb_order.Response{
				Error:  decreaseResp.Response.Error,
				Status: decreaseResp.Response.Status,
			},
		}, nil
	}

	return &pb_order.CreateOrderResponse{
		Response: &pb_order.Response{
			Status: http.StatusCreated,
		},
		Order: &pb_order.Order{
			OrderId:   order.Id,
			ProductId: order.ProductId,
			Quantity:  order.Quantity,
			UserId:    order.UserId,
		},
	}, nil
}
