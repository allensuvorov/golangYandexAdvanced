package main

import (
	"bytes"
	"crypto/md5"
	"crypto/rand"
	"fmt"
)

func main() {
	var (
		data  []byte         // слайс случайных байт
		hash1 []byte         // хеш с использованием интерфейса hash.Hash
		hash2 [md5.Size]byte // хеш, возвращаемый функцией md5.Sum
	)
	// допишите код
	// 1) генерация data длиной 512 байт
	data = make([]byte, 512)
	_, err := rand.Read(data) // записываем байты в массив data
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	// 2) вычисление hash1 с использованием md5.New
	// 3) вычисление hash2 функцией md5.Sum

	// ...

	// hash2[:] приводит массив байт к слайсу
	if bytes.Equal(hash1, hash2[:]) {
		fmt.Println("Всё правильно! Хеши равны")
	} else {
		fmt.Println("Что-то пошло не так")
	}
}
