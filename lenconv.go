package unitconv

type Length struct {
	FromValue float64
	FromUnit  int
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

func (l Length) GetToValue(toUnit int) (float64, error) {
	return l.base2To(l.from2Base(l.FromValue, l.FromUnit), toUnit), nil
}

func (l Length) from2Base(fromValue float64, fromUnit int) (baseValue float64) {
	switch fromUnit {
	case Meter:
		baseValue = fromValue
	case Kilometer:
		baseValue = fromValue * 1000
	case Centimeter:
		baseValue = fromValue / 100
	case Millimeter:
		baseValue = fromValue / 1000
	case Micrometer:
		baseValue = fromValue / 1000000
	case Nanometer:
		baseValue = fromValue / 1000000000
	case Mile:
		baseValue = fromValue * 1609.35
	case Yard:
		baseValue = fromValue * 0.9144
	case Foot:
		baseValue = fromValue * 0.3048
	case Inch:
		baseValue = fromValue * 0.0254
	case LightYear:
		baseValue = fromValue * 9.46066e15
	}
	return
}

func (l Length) base2To(baseValue float64, toUint int) (toValue float64) {
	switch toUint {
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
	}
	return
}
