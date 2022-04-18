package models

type Product struct {
	ProductId uint64 `json:"product_id,omitempty"`
	Name      string `json:"name,omitempty"`
	Sku       string `json:"sku,omitempty"`
	Stock     uint64 `json:"stock,omitempty"`
	Price     uint64 `json:"price,omitempty"`
}
