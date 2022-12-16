package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

var ErrProductNotFound = fmt.Errorf("Product not found")

type Product struct {
	ID          int     `json:"id" `
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Price       float32 `json:"price" binding:"required"`
	SKU         string  `json:"sku" binding:"required"`
	CreatedAt   string  `json:"-"`
	UpdatedAt   string  `json:"-"`
	DeletedAt   string  `json:"-"`
}

type Products []*Product

func GetProducts() Products {
	return productList
}
func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func FromJSON(r io.Reader) (*Product, error) {
	var prod *Product
	err := json.NewDecoder(r).Decode(&prod)
	if err != nil {
		return nil, err
	}
	return prod, nil
}

func getNextID() int {
	lp := productList[len(productList)-1]
	return lp.ID + 1
}

func AddProduct(prod *Product) {
	prod.ID = getNextID()
	productList = append(productList, prod)
}

func findProduct(id int) (*Product, int, error) {
	for i, p := range productList {
		if p.ID == id {
			return p, i, nil
		}
	}
	fmt.Println("Product not found")
	return nil, -1, ErrProductNotFound
}
func UpdateProduct(id int, prod *Product) error {
	_, pos, err := findProduct(id)
	if err != nil {
		return err
	}
	prod.ID = id
	productList[pos] = prod
	return nil
}

var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc323",
		CreatedAt:   time.Now().UTC().String(),
		UpdatedAt:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "fjd34",
		CreatedAt:   time.Now().UTC().String(),
		UpdatedAt:   time.Now().UTC().String(),
	},
}
