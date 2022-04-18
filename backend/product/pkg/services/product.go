package services

import (
	"context"
	"log"
	"net/http"

	"github.com/kbiits/microservices-grpc-simple-ecommerce-product-service/pkg/db"
	"github.com/kbiits/microservices-grpc-simple-ecommerce-product-service/pkg/models"
	"github.com/kbiits/microservices-grpc-simple-ecommerce-product-service/pkg/pb"
)

type Service struct {
	pb.UnimplementedProductServiceServer
	DB *db.DB
}

func (s *Service) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	product := &models.Product{
		Name:  req.Product.Name,
		Stock: req.Product.Stock,
		Price: req.Product.Price,
		Sku: req.Product.Sku,
	}

	if result := s.DB.Create(product); result.Error != nil {
		return &pb.CreateProductResponse{
			Response: &pb.Response{
				Error:  "Failed to create product",
				Status: http.StatusInternalServerError,
			},
		}, nil
	}

	return &pb.CreateProductResponse{
		Response: &pb.Response{
			Status: http.StatusCreated,
		},
		Product: &pb.Product{
			ProductId: product.ProductId,
			Name:      product.Name,
			Sku:       product.Sku,
			Stock:     product.Stock,
			Price:     product.Price,
		},
	}, nil
}

func (s *Service) FindOne(ctx context.Context, req *pb.FindOneRequest) (*pb.FindOneResponse, error) {
	product := &models.Product{}

	log.Println("FIND ONE Request")
	log.Println("product id :", req.ProductId)

	reqModel := &models.Product{ProductId: req.ProductId}
	log.Println("DB Query Condition", reqModel)
	if result := s.DB.First(product, req.ProductId); result.Error != nil {
		return &pb.FindOneResponse{
			Response: &pb.Response{
				Error:  "Product not found",
				Status: http.StatusNotFound,
			},
		}, nil
	}

	return &pb.FindOneResponse{
		Product: &pb.Product{
			ProductId: product.ProductId,
			Name:      product.Name,
			Sku:       product.Sku,
			Stock:     product.Stock,
			Price:     product.Price,
		},
		Response: &pb.Response{
			Status: http.StatusOK,
		},
	}, nil
}

func (s *Service) ListProducts(ctx context.Context, req *pb.ListProductsRequest) (*pb.ListProductsResponse, error) {
	var products []models.Product
	if result := s.DB.Find(&products); result.Error != nil {
		return &pb.ListProductsResponse{
			Response: &pb.Response{
				Error:  "Failed to get products",
				Status: http.StatusInternalServerError,
			},
		}, nil
	}

	var results []*pb.Product
	for _, p := range products {
		results = append(results, &pb.Product{
			ProductId: p.ProductId,
			Name:      p.Name,
			Sku:       p.Sku,
			Stock:     p.Stock,
			Price:     p.Price,
		})
	}

	return &pb.ListProductsResponse{
		Response: &pb.Response{
			Status: http.StatusOK,
		},
		Products: results,
	}, nil
}
func (s *Service) DecreaseStock(ctx context.Context, req *pb.DecreaseStockRequest) (*pb.DecreaseStockResponse, error) {
	var product models.Product

	if result := s.DB.First(&product, req.ProductId); result.Error != nil {
		return &pb.DecreaseStockResponse{
			Response: &pb.Response{
				Status: http.StatusNotFound,
				Error:  result.Error.Error(),
			},
		}, nil
	}

	if product.Stock < req.Quantity {
		return &pb.DecreaseStockResponse{
			Response: &pb.Response{
				Status: http.StatusConflict,
				Error:  "Stock too low",
			},
		}, nil
	}

	var log models.StockDecreaseLog

	if result := s.DB.Where(&models.StockDecreaseLog{OrderId: req.OrderId}).First(&log); result.Error == nil {
		return &pb.DecreaseStockResponse{
			Response: &pb.Response{
				Status: http.StatusConflict,
				Error:  "Stock already decreased",
			},
		}, nil
	}

	// decrease stock
	product.Stock = product.Stock - req.Quantity
	s.DB.Save(&product)

	// save logs
	log.OrderId = req.OrderId
	log.ProductRefer = product.ProductId
	s.DB.Create(&log)

	return &pb.DecreaseStockResponse{
		Response: &pb.Response{
			Status: http.StatusOK,
		},
	}, nil
}
