package golang10_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

/**
JSON Array

- Selain tipe dalam bentuk Object, biasanya dalam JSON, kita kadang menggunakan tipe data Array
- Array di JSON mirip dengan Array di JavaScript, dia bisa berisikan tipe data primitif, atau tipe data kompleks
  (Object atau Array)
- Di Go-Lang, JSON Array direpresentasikan dalam bentuk slice
- Konversi dari JSON atau ke JSON dilakukan secara otomatis oleh package json menggunakan tipe data slice
*/

type Customer2 struct {
	FirstName  string
	MiddleName string
	LastName   string
	Hobbies    []string
}

func TestJsonArray(t *testing.T) {
	customer := Customer2{
		FirstName:  "Adnin",
		MiddleName: "Rifandi",
		LastName:   "Sutanto",
		Hobbies:    []string{"Learning", "Coding", "Gaming"},
	}

	bytes, err := json.Marshal(customer)
	if err != nil {
		return
	}

	fmt.Println(string(bytes))
}

func TestJsonArrayDecode(t *testing.T) {
	jsonRequest := `{"FirstName":"Adnin","MiddleName":"Rifandi","LastName":"Sutanto","Hobbies":["Learning","Coding","Gaming"]}`
	jsonBytes := []byte(jsonRequest)

	customer := Customer2{}
	err := json.Unmarshal(jsonBytes, &customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
}

type Address struct {
	Street     string
	Country    string
	PostalCode string
}

type Customer3 struct {
	FirstName  string
	MiddleName string
	LastName   string
	Hobbies    []string
	Address    []Address
}

func TestJsonArrayObject(t *testing.T) {
	customer := Customer3{
		FirstName:  "Adnin",
		MiddleName: "Rifandi",
		LastName:   "Sutanto",
		Hobbies:    []string{"Learning", "Coding", "Gaming"},
		Address: []Address{
			{
				Street:     "Sinagar",
				Country:    "Indonesia",
				PostalCode: "16620",
			},
		},
	}

	bytes, err := json.Marshal(customer)
	if err != nil {
		return
	}

	fmt.Println(string(bytes))
}

func TestJsonArrayObjectDecode(t *testing.T) {
	jsonRequest := `{"FirstName":"Adnin","MiddleName":"Rifandi","LastName":"Sutanto","Hobbies":["Learning","Coding","Gaming"],"Address":[{"Street":"Sinagar","Country":"Indonesia","PostalCode":"16620"}]}`
	jsonBytes := []byte(jsonRequest)

	customer := Customer3{}
	err := json.Unmarshal(jsonBytes, &customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
}

func TestDecodeJsonArray(t *testing.T) {
	jsonArray := `[{"Street":"Sinagar","Country":"Indonesia","PostalCode":"16620"}]`
	jsonBytes := []byte(jsonArray)

	addresses := []Address{}
	err := json.Unmarshal(jsonBytes, &addresses)
	if err != nil {
		panic(err)
	}

	fmt.Println(addresses)
	fmt.Println(len(addresses))
}
