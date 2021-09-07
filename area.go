package unitconv

type Area struct {
	unitconv
}

const (
	SquareMeter = iota
	SquareKilometer
	SquareCentimeter
	SquareMillimeter
	SquareMicrometer
	Hectare
	SquareMile
	SquareYard
	SquareFoot
	SquareInch
	Acre
)

func NewArea(fromValue float64, fromUnit int) Unitconv {
	return &Area{unitconv{
		fromValue: fromValue,
		fromUnit:  fromUnit,
	}}
}

func (a *Area) GetToValue(toUnit int) (toValue float64, err error) {
	var baseValue float64
	switch a.fromUnit {
	case SquareMeter:
		baseValue = a.fromValue
	case SquareKilometer:
		baseValue = a.fromValue * 1000000
	case SquareCentimeter:
		baseValue = a.fromValue * 0.0001
	case SquareMillimeter:
		baseValue = a.fromValue * 0.000001
	case SquareMicrometer:
		baseValue = a.fromValue * 0.000000000001
	case Hectare:
		baseValue = a.fromValue * 10000
	case SquareMile:
		baseValue = a.fromValue * 2589990
	case SquareYard:
		baseValue = a.fromValue * 0.83612736
	case SquareFoot:
		baseValue = a.fromValue * 0.0929030
	case SquareInch:
		baseValue = a.fromValue * 0.000645160
	case Acre:
		baseValue = a.fromValue * 4046.8564224
	default:
		err = ErrInvalidUnit
	}

	switch toUnit {
	case SquareMeter:
		toValue = baseValue / 1
	case SquareKilometer:
		toValue = baseValue / 1000000
	case SquareCentimeter:
		toValue = baseValue / 0.0001
	case SquareMillimeter:
		toValue = baseValue / 0.000001
	case SquareMicrometer:
		toValue = baseValue / 0.000000000001
	case Hectare:
		toValue = baseValue / 10000
	case SquareMile:
		toValue = baseValue / 2589990
	case SquareYard:
		toValue = baseValue / 0.83612736
	case SquareFoot:
		toValue = baseValue / 0.0929030
	case SquareInch:
		toValue = baseValue / 0.000645160
	case Acre:
		toValue = baseValue / 4046.8564224
	default:
		err = ErrInvalidUnit
	}
	return
}
