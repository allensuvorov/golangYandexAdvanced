package main

import (
	"fmt"
	"log"

	"github.com/pelletier/go-toml"
	"gopkg.in/yaml.v2"
)

type Data struct {
	ID     int    `toml:"id"`
	Name   string `toml:"name"`
	Values []byte `toml:"values"`
}

const yamlData = `
id: 101
name: John Doe
values:
- 11
- 22
- 33
`

func main() {
	// допишите код
	// 1) десериализуйте yamlData в переменную типа Data
	d := Data{}
	err := yaml.Unmarshal([]byte(yamlData), &d)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	// 2) преобразуйте полученную переменную в TOML
	t, err := toml.Marshal(d)
	if err != nil {
		panic(err)
	}
	// 3) выведите в консоль результат
	// fmt.Println("data", d)
	fmt.Println("toml", string(t))
}
