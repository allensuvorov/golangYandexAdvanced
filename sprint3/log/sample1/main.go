package main

import (
	"log"
	"os"
)

func main() {
	flog, err := os.OpenFile(`server.log`, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer flog.Close()
	mylog := log.New(flog, `serv `, log.LstdFlags|log.Lshortfile)
	mylog.Println(`Start server`)
	mylog.Println(`Finish server`)
}
