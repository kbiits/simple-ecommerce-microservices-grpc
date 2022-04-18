package models

type Product struct {
	ProductId         uint64           `json:"id" gorm:"primaryKey"`
	Name              string           `json:"name"`
	Sku               string           `json:"sku"`
	Stock             uint64           `json:"stock"`
	Price             uint64           `json:"price"`
	StockDecreaseLogs StockDecreaseLog `gorm:"foreignKey:ProductRefer"`
}
