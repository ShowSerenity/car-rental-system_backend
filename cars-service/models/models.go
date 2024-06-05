package models

type Car struct {
	ID      int     `json:"id" gorm:"primaryKey"`
	Make    string  `json:"make,omitempty"`
	Model   string  `json:"model,omitempty"`
	Year    int     `json:"year,omitempty"`
	Color   string  `json:"color,omitempty"`
	Mileage int     `json:"mileage,omitempty"`
	Price   float64 `json:"price,omitempty"`
}
