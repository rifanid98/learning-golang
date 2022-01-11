package golang10_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

/**
Map

- Saat menggunakan JSON, kadang mungkin kita menemukan kasus data JSON nya dynamic
- Artinya atribut nya tidak menentu, bisa bertambah, bisa berkurang, dan tidak tetap
- Pada kasus seperti itu, menggunakan Struct akan menyulitkan, karena pada Struct, kita harus menentukan semua atributnya
- Untuk kasus seperti ini, kita bisa menggunakan tipe data map[string]interface{}
- Secara otomatis, atribut akan menjadi key di map, dan value menjadi value di map
- Namun karena value berupa interface{}, maka kita harus lakukan konversi secara manual jika ingin mengambil value nya
- Dan tipe data Map tidak mendukung JSON Tag lagi
*/

func TestMapDecode(t *testing.T) {
	jsonRequest := `{"id":"1","name":"Product 1","price":"10000","image_url":"product.png"}`
	jsonBytes := []byte(jsonRequest)

	var result map[string]interface{}
	err := json.Unmarshal(jsonBytes, &result)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
	fmt.Println(result["id"])
	fmt.Println(result["name"])
	fmt.Println(result["price"])
	fmt.Println(result["image_url"])
}

func TestMapEncode(t *testing.T) {
	product := map[string]interface{}{
		"id":        "1",
		"name":      "Product 1",
		"price":     "10000",
		"image_url": "product.png",
	}

	bytes, err := json.Marshal(product)
	if err != nil {
		return
	}

	fmt.Println(string(bytes))
}
