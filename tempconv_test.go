package unitconv

import (
	"math"
	"testing"
)

func TestTempConv(t *testing.T) {
	tests := []struct {
		fromValue float64
		fromUnit  int
		toUnit    int
		want      float64
	}{
		{0, Kelvin, Kelvin, 0},
		{0, Kelvin, Celsius, -273.15},
		{0, Kelvin, Fahrenheit, -459.67},
		{0, Celsius, Kelvin, 273.15},
		{0, Celsius, Celsius, 0},
		{0, Celsius, Fahrenheit, 32},
		{0, Fahrenheit, Kelvin, 255.37},
		{0, Fahrenheit, Celsius, -17.78},
		{0, Fahrenheit, Fahrenheit, 0},
	}

	for _, test := range tests {
		conv := NewTempConv(test.fromValue, test.fromUnit, test.toUnit)
		got := conv.To.Value
		if diff := math.Abs(got - test.want); diff > 0.01 {
			t.Errorf("got %f != want %f", got, test.want)
		}
	}
}
