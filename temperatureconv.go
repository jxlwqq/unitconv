package unitconv

type TemperatureConv struct {
	From
	To
}

const (
	Kelvin = iota // baseUnit
	Celsius
	Fahrenheit
)

func NewTemperatureConv(FromValue float64, FromUnit, toUnit int) *TemperatureConv {
	t := &TemperatureConv{
		From{Value: FromValue, Unit: FromUnit},
		To{Unit: toUnit},
	}

	t.conv()

	return t
}

func (t *TemperatureConv) conv() {
	var baseValue float64
	switch t.From.Unit {
	case Kelvin:
		baseValue = t.From.Value
	case Celsius:
		baseValue = t.From.Value + 273.15
	case Fahrenheit:
		baseValue = (t.From.Value + 459.67) * 5 / 9
	}

	switch t.To.Unit {
	case Kelvin:
		t.To.Value = baseValue
	case Celsius:
		t.To.Value = baseValue - 273.15
	case Fahrenheit:
		t.To.Value = baseValue*9/5 - 459.67
	}
}
