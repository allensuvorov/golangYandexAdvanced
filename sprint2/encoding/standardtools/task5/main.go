package main

import (
	"encoding/xml"
	"fmt"
	"log"
)

type Person struct {
	ID     int // ...
	Name   string
	Email  string
	Phones []string // ...
}

type List struct {
	Persons []Person // ...
}

func main() {
	var v List
	data := `
    <List>
        <Person id="1">
            <Name>Carla Mitchel</Name>
            <Phones>
                <Phone>123-45-67</Phone>
                <Phone>890-12-34</Phone>
            </Phones>
        </Person>
        <Person id="2">
            <Name>Michael Smith</Name>
            <Email>msmith@example.com</Email>
        </Person>
    </List>
    `
	err := xml.Unmarshal([]byte(data), &v)
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range v.Persons {
		fmt.Println(item.ID, item.Name, item.Email, item.Phones)
	}
}
