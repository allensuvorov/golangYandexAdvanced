package main

import (
	"time"
	"fmt"
	"errors"
)

type TimeError struct {
	Time time.Time
	Err error
}

func (te *TimeError) Error() string {
	return fmt.Sprintf("%v %v", te.Time.Format(`2006/01/02 15:04:05`), te.Err)
}

func NewTimeError(err error) error {
	return &TimeError{
		Time: time.Now(),
		Err: err,
	}
}

func main(){
	ne := errors.New("New Error")
	nte := NewTimeError(ne) 
	fmt.Printf("Types: %T, %T", ne, nte)
}
