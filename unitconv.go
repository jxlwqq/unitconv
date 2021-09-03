package unitconv

import "errors"

var ErrInvalidUnit = errors.New("invalid unit")

type Unitconv interface {
	GetToValue(toUnit int) (toValue float64, err error)
}

type unitconv struct {
	fromValue float64
	fromUnit  int
}
