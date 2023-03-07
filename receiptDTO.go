package main

type listItemDTO struct {
	ShortDescription string  `json:"shortDescription"`
	Price            float32 `json:"price,string"`
}

type receiptDTO struct {
	Retailer      string        `json:"retailer"`
	PurchasedDate string        `json:"purchaseDate"`
	PurchasedTime string        `json:"purchaseTime"`
	Items         []listItemDTO `json:"items"`
	Total         float32       `json:"total,string"`
}

type receiptCreateResponse struct {
	Id string `json:"id"`
}

type receiptGetPointsResponse struct {
	Points int `json:"points"`
}
