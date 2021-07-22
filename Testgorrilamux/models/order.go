package models

// type Order struct {
// 	Id        uint   `json:"id"`
// 	UserId    uint   `json:"user_id"`
// 	ProductId uint   `json:"product_id"`
// 	Quantity  uint   `json:"quantity"`
// 	Payment   string `json:"payment"`
// 	Total     uint   `json:"total"`
// }

type OrderItem struct {
	Id        uint     `json:"id"`
	UserId    uint     `json:"user_id"`
	ProductId uint     `json:"product_id"`
	Quantity  uint     `json:"quantity"`
	IsPaid    int      `json:"is_paid"`
	Product   Product  `json:"product"`
	User      Customer `json:"user"`
}
