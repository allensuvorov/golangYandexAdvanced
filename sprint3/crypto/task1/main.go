package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
)

func randB(l int) string {
	// определяем слайс нужной длины
	b := make([]byte, l)
	_, err := rand.Read(b) // записываем байты в массив b
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}
	return base64.StdEncoding.EncodeToString(b)
}

func main() {
	fmt.Println(randB(5))
}
