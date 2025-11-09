package models

type Product struct {
	ID                 int     `json:"id"`
	Title              string  `json:"name"`
	Description        string  `json:"description"`
	Rating             float64 `json:"rating"`
	*Category          `json:"category"`
	*Brand             `json:"brand"`
	Price              float64  `json:"price"`
	DiscountPercentage float64  `json:"discount_percentage"`
	Tags               []string `json:"tags"`
	Weight             int      `json:"weight"`
}
