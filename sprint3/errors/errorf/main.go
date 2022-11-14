package main

import( 
	"fmt"
	"errors"
)

type dog struct{}

func (d *dog) Error() string{
	return "good dog"
} 

func main(){
	ef := fmt.Errorf("test error %v", "fmt.Errorf")
	ne := errors.New("errors.New test")
	gs := error(&dog{})
	fmt.Printf("Type is: %T, %T, %T \n", ef, ne, gs)
}

