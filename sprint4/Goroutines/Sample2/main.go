package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Printf("Не копируем переменную i: %d\n", i)
		}()
	}
	time.Sleep(1 * time.Second)

	for i := 0; i < 5; i++ {
		go func(v int) {
			fmt.Printf("Копируем переменную i: %d\n", v)
		}(i)
	}
	time.Sleep(1 * time.Second)
}
