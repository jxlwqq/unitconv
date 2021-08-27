package unitconv

type Length struct {
	fromValue float64
	formUnit  int
}

const (
	Meter = iota // base unit
	Kilometer
	Centimeter
	Millimeter
)

func (l *Length) GetToValue(toUnit int) (float64, error) {
	// todo
	return l.fromValue, nil
}
