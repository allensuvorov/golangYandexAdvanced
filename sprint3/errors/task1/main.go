package main

import (
	"fmt"
	"os"
	"strings"
	//"strings"
)

// LabelError описывает ошибку с дополнительной меткой.
type LabelError struct {
	Label string // метка должна быть в верхнем регистре
	Err   error
}

// добавьте методы Error() и NewLabelError(string, error)
func (le *LabelError) Error() string {
	return fmt.Sprintf("[%v] %v", strings.ToUpper(le.Label), le.Err)
}

func NewLabelError(s string, err error) error {
	return &LabelError{
		Label: s,
		Err:   err,
	}
}

func main() {
	_, err := os.ReadFile("mytest.txt")
	if err != nil {
		err = NewLabelError("file", err)
	}
	fmt.Println(err)
	// должна выводить
	// [FILE] open mytest.txt: no such file or directory
}
