// Андрей родился 26 ноября 1993 года. Посчитайте количество дней до его 100-летия — относительно сегодняшнего дня.
package main

import (
	"fmt"
	"time"
)

func main() {
	// скопируйте блок себе в IDE и допишите код
	birthday := time.Date(2093, time.November, 26, 0, 0, 0, 0, time.Local)
	days := birthday.Sub(time.Now()).Truncate(time.Hour*24).Hours() / 24
	fmt.Println(days)
}
