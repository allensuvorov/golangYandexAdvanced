package main

import (
	"bytes"
	"fmt"
	"log"
)

func main() {
	var buf bytes.Buffer
	logger := log.New(&buf, "mylog: ", 0)
	logger.Print("Hello, world!")
	logger.Print("Goodbye")
	fmt.Print(&buf)	
}
