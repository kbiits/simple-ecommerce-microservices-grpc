package models

type StockDecreaseLog struct {
	Id           uint64 `json:"id" gorm:"primaryKey"`
    OrderId      uint64 `json:"order_id"`
    ProductRefer uint64 `json:"product_id"`
}