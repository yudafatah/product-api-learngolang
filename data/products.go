package data

import (
	"encoding/json"
	"io"
	"time"
)

//Product structure defination
type Product struct {
	ID 				int `json:"id"`
	Name 			string `json:"name"`
	Description 	string `json:"description"`
	Price 			float32 `json:"price"`
	SKU 			string `json:"sku"`
	CreatedOn 		string `json:"createdOn"`
	UpdatedOn 		string `json:"updatedOn"`
	DeletedOn 		string `json:"deletedOn"`
}

type Products []*Product 

//ToJSON encode json and return encoded json
func (p*Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

//GetProduct : get list of product 
func GetProduct() Products {
	return productList
}

var productList = []*Product{
	{
		ID:          1,
		Name:        "Latte",
		Description: "des 1",
		Price:       2.22,
		SKU:         "sku1",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
		DeletedOn:    "",
	},
	{
		ID:          2,
		Name:        "Caramel",
		Description: "des 2",
		Price:       2.44,
		SKU:         "sku2",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
		DeletedOn:    "",
	},
}