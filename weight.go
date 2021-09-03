package unitconv

const (
	Kilogram = iota
	Gram
	Milligram
	MetricTon
	LongTon
	ShortTon
	Pound
	Ounce
	Carrat
	AtomicMassUnit
)

type weight struct {
	unitconv
}

func NewWeight(fromValue float64, fromUnit int) Unitconv {
	return &weight{
		unitconv{
			fromValue: fromValue,
			fromUnit:  fromUnit,
		},
	}

}

func (w *weight) GetToValue(toUnit int) (toValue float64, err error) {
	var baseValue float64
	switch w.fromUnit {
	case Kilogram:
		baseValue = w.fromValue
	case Gram:
		baseValue = w.fromValue / 1000
	case Milligram:
		baseValue = w.fromValue / 1000000
	case MetricTon:
		baseValue = w.fromValue * 1000
	case LongTon:
		baseValue = w.fromValue * 1016.04608
	case ShortTon:
		baseValue = w.fromValue * 907.184
	case Pound:
		baseValue = w.fromValue * 0.453592
	case Ounce:
		baseValue = w.fromValue * 0.0283495
	case Carrat:
		baseValue = w.fromValue * 0.0002
	case AtomicMassUnit:
		baseValue = w.fromValue / 6.022136652e+26
	default:
		err = ErrInvalidUnit
	}

	switch toUnit {
	case Kilogram:
		toValue = baseValue
	case Gram:
		toValue = baseValue * 1000
	case Milligram:
		toValue = baseValue * 1000000
	case MetricTon:
		toValue = baseValue / 1000
	case LongTon:
		toValue = baseValue / 1016.04608
	case ShortTon:
		toValue = baseValue / 907.184
	case Pound:
		toValue = baseValue / 0.453592
	case Ounce:
		toValue = baseValue / 0.0283495
	case Carrat:
		toValue = baseValue / 0.0002
	case AtomicMassUnit:
		toValue = baseValue * 6.022136652e+26
	default:
		err = ErrInvalidUnit
	}

	return
}
