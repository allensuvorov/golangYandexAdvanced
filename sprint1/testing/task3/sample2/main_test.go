package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFullName(t *testing.T) {
	tests := []struct { // добавился слайс тестов
		name  string
		value User
		want  string
	}{
		{
			name:  "empty first name", // описывается каждый тест
			value: User{"", "Popov"},  // значения, которые будет принимать функция
			want:  "Popov",            // ожидаемое значение
		},
		{
			name:  "empty last name",
			value: User{"Sasha", ""},
			want:  "Sasha",
		},
		{
			name:  "all caps",
			value: User{"SASHA", "IVANOV"},
			want:  "SASHA IVANOV",
		},
		{
			name:  "with middle name",
			value: User{"Ivan Ivanovich", "Ivanov"},
			want:  "Ivan Ivanovich Ivanov",
		},
		{
			name:  "all lower case",
			value: User{"ivan", "ivanov"},
			want:  "ivan ivanov",
		},
	}
	for _, tt := range tests { // цикл по всем тестам
		t.Run(tt.name, func(t *testing.T) {
			v := User.FullName(tt.value)
			assert.Equal(t, tt.want, v, "values should be equal")
		})
	}
}
