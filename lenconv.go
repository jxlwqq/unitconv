package unitconv

type LenConv struct {
	From
	To
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

func NewLenConv(FromValue float64, FromUnit, toUnit int) *LenConv {
	l := &LenConv{
		From{Value: FromValue, Unit: FromUnit},
		To{Unit: toUnit},
	}

	l.conv()

	return l
}

func (l *LenConv) conv() {
	var baseValue float64
	switch l.From.Unit {
	case Meter:
		baseValue = l.From.Value
	case Kilometer:
		baseValue = l.From.Value * 1000
	case Centimeter:
		baseValue = l.From.Value / 100
	case Millimeter:
		baseValue = l.From.Value / 1000
	case Micrometer:
		baseValue = l.From.Value / 1000000
	case Nanometer:
		baseValue = l.From.Value / 1000000000
	case Mile:
		baseValue = l.From.Value * 1609.35
	case Yard:
		baseValue = l.From.Value * 0.9144
	case Foot:
		baseValue = l.From.Value * 0.3048
	case Inch:
		baseValue = l.From.Value * 0.0254
	case LightYear:
		baseValue = l.From.Value * 9.46066e15
	}

	switch l.To.Unit {
	case Meter:
		l.To.Value = baseValue
	case Kilometer:
		l.To.Value = baseValue / 1000
	case Centimeter:
		l.To.Value = baseValue * 100
	case Millimeter:
		l.To.Value = baseValue * 1000
	case Micrometer:
		l.To.Value = baseValue * 1000000
	case Nanometer:
		l.To.Value = baseValue * 1000000000
	case Mile:
		l.To.Value = baseValue / 1609.35
	case Yard:
		l.To.Value = baseValue / 0.9144
	case Foot:
		l.To.Value = baseValue / 0.3048
	case Inch:
		l.To.Value = baseValue / 0.0254
	case LightYear:
		l.To.Value = baseValue / 9.46066e15
	}

}
