package unitconv

const (
	Kelvin = iota // baseUnit
	Celsius
	Fahrenheit
)

type temperature struct {
	unitconv
}

func NewTemperature(fromValue float64, fromUnit int) Unitconv {
	return &temperature{
		unitconv{
			fromValue: fromValue,
			fromUnit:  fromUnit,
		},
	}
}

func (t *temperature) GetToValue(toUnit int) (toValue float64, err error) {
	var baseValue float64
	switch t.fromUnit {
	case Kelvin:
		baseValue = t.fromValue
	case Celsius:
		baseValue = t.fromValue + 273.15
	case Fahrenheit:
		baseValue = (t.fromValue + 459.67) * 5 / 9
	default:
		err = InvalidUnitErr
	}

	switch toUnit {
	case Kelvin:
		toValue = baseValue
	case Celsius:
		toValue = baseValue - 273.15
	case Fahrenheit:
		toValue = baseValue*9/5 - 459.67
	default:
		err = InvalidUnitErr
	}

	return
}
