package main

// import (
//     "fmt"

//     "github.com/pelletier/go-toml"
//     "gopkg.in/yaml.v2"
// )

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
	// 2) преобразуйте полученную переменную в TOML
	// 3) выведите в консоль результат
}
