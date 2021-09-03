package unitconv

const (
	Second = iota // baseUnit
	Millisecond
	Microsecond
	Nanosecond
	Picosecond
	Minute
	Hour
	Day
	Week
	Month
	Year
)

type time struct {
	unitconv
}

func NewTime(fromValue float64, fromUnit int) Unitconv {
	return &time{unitconv{
		fromValue: fromValue,
		fromUnit:  fromUnit,
	}}
}

func (t *time) GetToValue(toUnit int) (toValue float64, err error) {
	var baseValue float64
	switch t.fromUnit {
	case Second:
		baseValue = t.fromValue
	case Millisecond:
		baseValue = t.fromValue * 0.001
	case Microsecond:
		baseValue = t.fromValue * 0.000001
	case Nanosecond:
		baseValue = t.fromValue * 0.000000001
	case Picosecond:
		baseValue = t.fromValue * 0.000000000001
	case Minute:
		baseValue = t.fromValue * 60
	case Hour:
		baseValue = t.fromValue * 3600
	case Day:
		baseValue = t.fromValue * 86400
	case Week:
		baseValue = t.fromValue * 604800
	case Month:
		baseValue = t.fromValue * 2629800
	case Year:
		baseValue = t.fromValue * 31557600
	default:
		err = InvalidUnitErr
	}

	switch toUnit {
	case Second:
		toValue = baseValue
	case Millisecond:
		toValue = baseValue / 0.001
	case Microsecond:
		toValue = baseValue / 0.000001
	case Nanosecond:
		toValue = baseValue / 0.000000001
	case Picosecond:
		toValue = baseValue / 0.000000000001
	case Minute:
		toValue = baseValue / 60
	case Hour:
		toValue = baseValue / 3600
	case Day:
		toValue = baseValue / 86400
	case Week:
		toValue = baseValue / 604800
	case Month:
		toValue = baseValue / 2629800
	case Year:
		toValue = baseValue / 31557600
	default:
		err = InvalidUnitErr
	}

	return
}
