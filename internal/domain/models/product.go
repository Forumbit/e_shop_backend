package models

type Product struct {
	ID          int      `json:"id"`
	Title       string   `json:"name"`
	Description string   `json:"description"`
	Rating      float64  `json:"rating"`
	CategoryID  *int     `json:"category"`
	BrandID     *int     `json:"brand"`
	Price       float64  `json:"price"`
	Discount    float64  `json:"discount"`
	Tags        []string `json:"tags"`
	Weight      int      `json:"weight"`
}
