package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		fmt.Print(s + ` `)
		time.Sleep(50 * time.Millisecond)
	}
}

func main() {
	go say(`hello`)
	go func() {
		say(`world`)
	}()
	say(`bye`)
}
