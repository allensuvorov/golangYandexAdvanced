package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Data struct {
	// определите здесь нужные поля
	ID      int    `json:"-"`
	Name    string `json:"name,omitempty"`
	Company string `json:"company,omitempty"`
}

func main() {
	foo := []Data{
		{
			ID:   10,
			Name: "John Doe",
		},
		{
			Name:    "Вася",
			Company: "Яндекс",
		},
	}
	out, err := json.Marshal(foo)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(out))
}
