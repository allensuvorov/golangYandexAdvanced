package main

import (
	"fmt"
	"os"
)
{}

func ReadTextFile(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return ``, fmt.Errorf(`failed reading file: %w`, err)
	}
	return string(data), nil
}

func main() {
	if data, err := ReadTextFile(`myconfig.yaml`); err != nil {
		if os.IsNotExist(error.Unwrap(err)) {
			// create config file
			return
		}
		panic(err)
	}
}