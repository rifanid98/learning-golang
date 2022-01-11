package golang10_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

/**
Decode JSON

- Sekarang kita sudah tahu bagaimana caranya melakukan encode dari tipe data di Go-Lang ke JSON
- Namun bagaimana jika kebalikannya?
- Untuk melakukan konversi dari JSON ke tipe data di Go-Lang (Decode), kita bisa menggunakan function
  json.Unmarshal(byte[], interface{})
- Dimana byte[] adalah data JSON nya, sedangkan interface{} adalah tempat menyimpan hasil konversi, biasa berupa pointer
*/

func TestDecodeJSON(t *testing.T) {
	jsonRequest := `{"FirstName":"Adnin","MiddleName":"Rifandi","LastName":"Sutanto"}`
	jsonBytes := []byte(jsonRequest)

	customer := Customer{}
	err := json.Unmarshal(jsonBytes, &customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
}
