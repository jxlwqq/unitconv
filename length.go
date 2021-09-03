package unitconv

type length struct {
	unitconv
}

const (
	Meter = iota // base unit
	Kilometer
	Centimeter
	Millimeter
	Micrometer
	Nanometer
	Mile
	Yard
	Foot
	Inch
	LightYear
)

func NewLength(fromValue float64, fromUnit int) Unitconv {
	return &length{
		unitconv{
			fromValue: fromValue,
			fromUnit:  fromUnit,
		},
	}
}

func (l *length) GetToValue(toUnit int) (toValue float64, err error) {
	var baseValue float64
	switch l.fromUnit {
	case Meter:
		baseValue = l.fromValue
	case Kilometer:
		baseValue = l.fromValue * 1000
	case Centimeter:
		baseValue = l.fromValue / 100
	case Millimeter:
		baseValue = l.fromValue / 1000
	case Micrometer:
		baseValue = l.fromValue / 1000000
	case Nanometer:
		baseValue = l.fromValue / 1000000000
	case Mile:
		baseValue = l.fromValue * 1609.35
	case Yard:
		baseValue = l.fromValue * 0.9144
	case Foot:
		baseValue = l.fromValue * 0.3048
	case Inch:
		baseValue = l.fromValue * 0.0254
	case LightYear:
		baseValue = l.fromValue * 9.46066e15
	default:
		err = ErrInvalidUnit
	}

	switch toUnit {
	case Meter:
		toValue = baseValue
	case Kilometer:
		toValue = baseValue / 1000
	case Centimeter:
		toValue = baseValue * 100
	case Millimeter:
		toValue = baseValue * 1000
	case Micrometer:
		toValue = baseValue * 1000000
	case Nanometer:
		toValue = baseValue * 1000000000
	case Mile:
		toValue = baseValue / 1609.35
	case Yard:
		toValue = baseValue / 0.9144
	case Foot:
		toValue = baseValue / 0.3048
	case Inch:
		toValue = baseValue / 0.0254
	case LightYear:
		toValue = baseValue / 9.46066e15
	default:
		err = ErrInvalidUnit
	}

	return
}
