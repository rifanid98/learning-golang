package golang10_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

/**
JSON Tag

- Secara default atribut yang terdapat di Struct dan JSON akan di mapping sesuai dengan nama atribut  yang sama
  (case sensitive)
- Kadang ada style yang berbeda antara penamaan atribute di Struct dan di JSON, misal di JSON kita ingin menggunakan
  snake_case, tapi di Struct, kita ingin menggunakan PascalCase
- Untungnya, package json mendukun Tag Reflection
- Kita bisa menambahkan tag reflection dengan nama json, lalu diikuti dengan atribut yang kita inginkan ketika konversi
  dari atau ke JSON
*/

type Product struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Price    string `json:"price"`
	ImageUrl string `json:"image_url"`
}

func TestJsonTag(t *testing.T) {
	product := Product{
		Id:       "1",
		Name:     "Product 1",
		Price:    "10000",
		ImageUrl: "product.png",
	}

	bytes, err := json.Marshal(product)
	if err != nil {
		return
	}

	fmt.Println(string(bytes))
}

func TestJsonTagDecode(t *testing.T) {
	jsonRequest := `{"id":"1","name":"Product 1","price":"10000","image_url":"product.png"}`
	jsonBytes := []byte(jsonRequest)

	product := Product{}
	err := json.Unmarshal(jsonBytes, &product)
	if err != nil {
		panic(err)
	}

	fmt.Println(product)
}
