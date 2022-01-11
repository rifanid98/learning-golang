package golang10_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

/**
Encode JSON

- Go-Lang telah menyediakan function untuk melakukan konversi data ke JSON, yaitu menggunakan function
  json.Marshal(interface{})
- Karena parameter nya adalah interface{}, maka kita bisa masukan tipe data apapun ke dalam function Marshal
*/

func logJson(data interface{}) {
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

func TestEncodeJson(t *testing.T) {
	logJson("Eko")
	logJson(1)
	logJson(true)
	logJson([]string{"Adnin", "Rifandi", "Sutanto", "Putra"})
}
