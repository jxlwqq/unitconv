package unitconv

type Temperature struct {
	FromValue float64
	FromUnit  int
}

const (
	Kelvin = iota // base unit
	Celsius
	Fahrenheit
)

func (t Temperature) GetToValue(toUnit int) (float64, error) {
	return t.base2To(t.from2Base(t.FromValue, t.FromUnit), toUnit), nil
}

func (t Temperature) from2Base(fromValue float64, fromUnit int) (baseValue float64) {
	switch fromUnit {
	case Kelvin:
		baseValue = fromValue
	case Celsius:
		baseValue = fromValue + 273.15
	case Fahrenheit:
		baseValue = (fromValue + 459.67) * 5 / 9
	}
	return
}

func (t Temperature) base2To(baseValue float64, toUint int) (toValue float64) {
	switch toUint {
	case Kelvin:
		toValue = baseValue
	case Celsius:
		toValue = baseValue - 273.15
	case Fahrenheit:
		toValue = baseValue*9/5 - 459.67
	}
	return
}
