package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	currentTimeStr := "2021-09-19T15:59:41+03:00"

	// we can use this
	layout := "2006-01-02T15:04:05-07:00"
	// or predefined layouts
	// layout := time.RFC3339 //	RFC3339     = "2006-01-02T15:04:05Z07:00"
	currentTime, err := time.Parse(layout, currentTimeStr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(currentTime)
}
