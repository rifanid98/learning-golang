package golang10_json

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

/**
Streaming Encoder

- Selain decoder, package json juga mendukung membuat Encoder yang bisa digunakan untuk menulis langsung JSON nya ke
  io.Writer
- Dengan begitu, kita tidak perlu menyimpan JSON datanya terlebih dahulu ke dalam variable string atau []byte, kita bisa
  langsung tulis ke io.Writer

json.Encoder

- Untuk membuat Encoder, kita bisa menggunakan function json.NewEncoder(writer)
- Dan untuk menulis data sebagai JSON langsung ke writer, kita bisa gunakan function Encode(interface{})
*/

func TestEncoder(t *testing.T) {
	writer, _ := os.Create("product_sample.json")
	encoder := json.NewEncoder(writer)

	product := Product{
		Id:       "1",
		Name:     "Product 1",
		Price:    "10000",
		ImageUrl: "product.png",
	}

	err := encoder.Encode(product)
	if err != nil {
		panic(err)
	}

	// verify created file product_sample.json
	reader, _ := os.Open("product_sample.json")
	decoder := json.NewDecoder(reader)

	product2 := Product{}
	err = decoder.Decode(&product2)
	if err != nil {
		panic(err)
	}

	fmt.Println(product2)
}
