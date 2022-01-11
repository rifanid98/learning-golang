package golang10_json

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

/**
Streaming Decoder

- Sebelumnya kita belajar package json dengan melakukan konversi data JSON yang sudah dalam bentuk variable dan data
  string atau []byte
- Pada kenyataanya, kadang data JSON nya berasal dari Input berupa io.Reader (File, Network, Request Body)
- Kita bisa saja membaca semua datanya terlebih dahulu, lalu simpan di variable, baru lakukan konversi dari JSON, namun
  hal ini sebenarnya tidak perlu dilakukan, karena package json memiliki fitur untuk membaca data dari Stream

json.Decoder

- Untuk membuat json Decoder, kita bisa menggunakan function json.NewDecoder(reader)
- Selanjutnya untuk membaca isi input reader dan konversikan secara langsung ke data di Go-Lang, cukup gunakan function
  Decode(interface{})
*/

func TestDecoder(t *testing.T) {
	reader, _ := os.Open("product.json")
	decoder := json.NewDecoder(reader)

	product := Product{}
	err := decoder.Decode(&product)
	if err != nil {
		panic(err)
	}

	fmt.Println(product)
}
