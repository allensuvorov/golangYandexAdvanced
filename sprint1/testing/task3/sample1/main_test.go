// для функции Abs проверьте тесты на -3, 3, -2.000001, -0.000000003 и другие значения.
package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAbs(t *testing.T) {
	tests := []struct { // добавился слайс тестов
		name  string
		value float64
		want  float64
	}{
		{
			name:  "negative int", // описывается каждый тест
			value: -3,             // значения, которые будет принимать функция
			want:  3,              // ожидаемое значение
		},
		{
			name:  "positive int",
			value: 3,
			want:  3,
		},
		{
			name:  "negative float",
			value: -2.000001,
			want:  2.000001,
		},
		{
			name:  "negative float with leading zero",
			value: -0.000000003,
			want:  0.000000003,
		},
		{
			name:  "negative big int - 100billion",
			value: -100_000_000_000,
			want:  100_000_000_000,
		},
	}
	for _, tt := range tests { // цикл по всем тестам
		t.Run(tt.name, func(t *testing.T) {
			v := Abs(tt.value)
			assert.Equal(t, tt.want, v, "values must be equal") // меняем на функцию Equal из пакета assert

		})
	}
}
