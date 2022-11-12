package main

import (
	"errors"
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
	return fmt.Sprintf("[%s] %v", le.Label, le.Err)
}

// Unwrap возвращает исходную ошибку.
func (le *LabelError) Unwrap() error {
	return le.Err
}

func NewLabelError(s string, err error) error {
	return &LabelError{
		Label: strings.ToUpper(s),
		Err:   err,
	}
}

func main() {
	_, err := os.ReadFile("mytest.txt")
	if err != nil {
		err = NewLabelError("file", err)
	}
	fmt.Println(errors.Is(err, os.ErrNotExist), err)
	// должна выводить
	// true [FILE] open mytest.txt: no such file or directory
}
