package models

type Order struct {
	ProductId uint64 `json:"product_id,omitempty"`
	Quantity  uint64 `json:"quantity,omitempty"`
	UserId    uint64 `json:"user_id,omitempty"`
	OrderId   uint64 `json:"order_id,omitempty"`
}
