package models

type Product struct {
	ID          int
	Name        string
	Price       float64
	TrademarkID *int
}

type ProductMap map[int]Product

type ProductList []Product
