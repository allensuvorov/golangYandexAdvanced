// Задание 8
// Используя Ticker, напишите программу, которая каждые 2 секунды 10 раз подряд выводит разницу в секундах между текущим временем и временем запуска программы. Выводить нужно только целую часть секунд.

package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	ticker := time.NewTicker(2 * time.Second) // created a ticker
	defer ticker.Stop()
	for i := 0; i < 10; i++ {
		t := <-ticker.C // wait for ticker
		fmt.Println(int(t.Sub(start).Seconds()))
	}
}
