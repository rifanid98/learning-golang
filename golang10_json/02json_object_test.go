package golang10_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

/**
JSON Object

- Pada materi sebelumnya kita melakukan encode data seperti string, number, boolean, dan tipe data primitif lainnya
- Walaupun memang bisa dilakukan, karena sesuai dengan kontrak interface{}, namun tidak sesuai dengan kontrak JSON
- Jika mengikuti kontrak json.org, data JSON bentuknya adalah Object dan Array
- Sedangkan value nya baru berupa

Struct

- JSON Object di Go-Lang direpresentasikan dengan tipe data Struct
- Dimana tiap attribute di JSON Object merupakan attribute di Struct
*/

type Customer struct {
	FirstName  string
	MiddleName string
	LastName   string
}

func TestJsonObject(t *testing.T) {
	customer := Customer{
		FirstName:  "Adnin",
		MiddleName: "Rifandi",
		LastName:   "Sutanto",
	}

	bytes, err := json.Marshal(customer)
	if err != nil {
		return
	}

	fmt.Println(string(bytes))
}
