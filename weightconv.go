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

type WeightConv struct {
	From
	To
}

func NewWeightConv(fromValue float64, fromUnit, toUnit int) *WeightConv {
	w := &WeightConv{
		From: From{Value: fromValue, Unit: fromUnit},
		To:   To{Unit: toUnit},
	}

	w.conv()

	return w
}

func (w *WeightConv) conv() {
	var baseValue float64
	switch w.From.Unit {
	case Kilogram:
		baseValue = w.From.Value
	case Gram:
		baseValue = w.From.Value / 1000
	case Milligram:
		baseValue = w.From.Value / 1000000
	case MetricTon:
		baseValue = w.From.Value * 1000
	case LongTon:
		baseValue = w.From.Value * 1016.04608
	case ShortTon:
		baseValue = w.From.Value * 907.184
	case Pound:
		baseValue = w.From.Value * 0.453592
	case Ounce:
		baseValue = w.From.Value * 0.0283495
	case Carrat:
		baseValue = w.From.Value * 0.0002
	case AtomicMassUnit:
		baseValue = w.From.Value / 6.022136652e+26
	}

	switch w.To.Unit {
	case Kilogram:
		w.To.Value = baseValue
	case Gram:
		w.To.Value = baseValue * 1000
	case Milligram:
		w.To.Value = baseValue * 1000000
	case MetricTon:
		w.To.Value = baseValue / 1000
	case LongTon:
		w.To.Value = baseValue / 1016.04608
	case ShortTon:
		w.To.Value = baseValue / 907.184
	case Pound:
		w.To.Value = baseValue / 0.453592
	case Ounce:
		w.To.Value = baseValue / 0.0283495
	case Carrat:
		w.To.Value = baseValue / 0.0002
	case AtomicMassUnit:
		w.To.Value = baseValue * 6.022136652e+26
	}
}
