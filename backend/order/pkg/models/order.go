package models

type Order struct {
	Id        uint64 `json:"id" gorm:"primaryKey"`
	Price     uint64 `json:"price"`
	ProductId uint64 `json:"product_id"`
	Quantity  uint64 `json:"quantity"`
	UserId    uint64 `json:"user_id"`
}
