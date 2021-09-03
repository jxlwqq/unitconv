package unitconv

import (
	"math"
	"testing"
)

func TestNewWeight(t *testing.T) {
	tests := []struct {
		fromValue float64
		fromUnit  int
		toUnit    int
		want      float64
	}{
		{1, Kilogram, Kilogram, 1},
		{1, Kilogram, Gram, 1000},
		{1, Kilogram, Milligram, 1000000},
		{1, Kilogram, MetricTon, 0.001},
		{1, Kilogram, LongTon, 0.0009842073},
		{1, Kilogram, ShortTon, 0.0011023122},
		{1, Kilogram, Pound, 2.2046244202},
		{1, Kilogram, Ounce, 35.273990723},
		{1, Kilogram, Carrat, 5000},
		{1, Kilogram, AtomicMassUnit, 6.022136652e+26},
		{1, Gram, Kilogram, 0.001},
		{1, Gram, Gram, 1},
		{1, Gram, Milligram, 1000},
		{1, Gram, MetricTon, 0.000001},
		{1, Gram, LongTon, 9.842073304e-7},
		{1, Gram, ShortTon, 0.0000011023},
		{1, Gram, Pound, 0.0022046244},
		{1, Gram, Ounce, 0.0352739907},
		{1, Gram, Carrat, 5},
		{1, Gram, AtomicMassUnit, 6.022136652e+23},
		{1, Milligram, Kilogram, 0.000001},
		{1, Milligram, Gram, 0.001},
		{1, Milligram, Milligram, 1},
		{1, Milligram, MetricTon, 1e-9},
		{1, Milligram, LongTon, 9.842073304e-10},
		{1, Milligram, ShortTon, 1.10231221e-9},
		{1, Milligram, Pound, 0.0000022046},
		{1, Milligram, Ounce, 0.000035274},
		{1, Milligram, Carrat, 0.005},
		{1, Milligram, AtomicMassUnit, 602213665200000000000},
		{1, MetricTon, Kilogram, 1000},
		{1, MetricTon, Gram, 1000000},
		{1, MetricTon, Milligram, 1000000000},
		{1, MetricTon, MetricTon, 1},
		{1, MetricTon, LongTon, 0.9842073304},
		{1, MetricTon, ShortTon, 1.1023122101},
		{1, MetricTon, Pound, 2204.6244202},
		{1, MetricTon, Ounce, 35273.990723},
		{1, MetricTon, Carrat, 5000000},
		{1, MetricTon, AtomicMassUnit, 6.022136652e+29},
		{1, LongTon, Kilogram, 1016.04608},
		{1, LongTon, Gram, 1016046.08},
		{1, LongTon, Milligram, 1016046080},
		{1, LongTon, MetricTon, 1.01604608},
		{1, LongTon, LongTon, 1},
		{1, LongTon, ShortTon, 1.12},
		{1, LongTon, Pound, 2240},
		{1, LongTon, Ounce, 35840},
		{1, LongTon, Carrat, 5080230.4},
		{1, LongTon, AtomicMassUnit, 611876833848892431306361667584},
		{1, ShortTon, Kilogram, 907.184},
		{1, ShortTon, Gram, 907184},
		{1, ShortTon, Milligram, 907184000},
		{1, ShortTon, MetricTon, 0.907184},
		{1, ShortTon, LongTon, 0.8928571429},
		{1, ShortTon, ShortTon, 1},
		{1, ShortTon, Pound, 2000},
		{1, ShortTon, Ounce, 32000},
		{1, ShortTon, Carrat, 4535920},
		{1, ShortTon, AtomicMassUnit, 546318601650796808640055476224},
		{1, Pound, Kilogram, 0.453592},
		{1, Pound, Gram, 453.592},
		{1, Pound, Milligram, 453592},
		{1, Pound, MetricTon, 0.000453592},
		{1, Pound, LongTon, 0.0004464286},
		{1, Pound, ShortTon, 0.0005},
		{1, Pound, Pound, 1},
		{1, Pound, Ounce, 16},
		{1, Pound, Carrat, 2267.96},
		{1, Pound, AtomicMassUnit, 273159300825398404869783552},
		{1, Ounce, Kilogram, 0.0283495},
		{1, Ounce, Gram, 28.3495},
		{1, Ounce, Milligram, 28349.5},
		{1, Ounce, MetricTon, 0.0000283495},
		{1, Ounce, LongTon, 0.0000279018},
		{1, Ounce, ShortTon, 0.00003125},
		{1, Ounce, Pound, 0.0625},
		{1, Ounce, Ounce, 1},
		{1, Ounce, Carrat, 141.7475},
		{1, Ounce, AtomicMassUnit, 17072456301587400304361472},
		{1, Carrat, Kilogram, 0.0002},
		{1, Carrat, Gram, 0.2},
		{1, Carrat, Milligram, 200},
		{1, Carrat, MetricTon, 2.e-7},
		{1, Carrat, LongTon, 1.96841466e-7},
		{1, Carrat, ShortTon, 2.20462442e-7},
		{1, Carrat, Pound, 0.0004409249},
		{1, Carrat, Ounce, 0.0070547981},
		{1, Carrat, Carrat, 1},
		{1, Carrat, AtomicMassUnit, 120442733040000003211264},
		{1, AtomicMassUnit, Kilogram, 1.660540199e-27},
		{1, AtomicMassUnit, Gram, 1.660540199e-24},
		{1, AtomicMassUnit, Milligram, 1.660540199e-21},
		{1, AtomicMassUnit, MetricTon, 1.660540199e-30},
		{1, AtomicMassUnit, LongTon, 1.634315837e-30},
		{1, AtomicMassUnit, ShortTon, 1.830433737e-30},
		{1, AtomicMassUnit, Pound, 3.660867475e-27},
		{1, AtomicMassUnit, Ounce, 5.85738796e-26},
		{1, AtomicMassUnit, Carrat, 8.302700999e-24},
		{1, AtomicMassUnit, AtomicMassUnit, 1},
	}

	for _, test := range tests {
		conv := NewWeight(test.fromValue, test.fromUnit)
		got, _ := conv.GetToValue(test.toUnit)
		if diff := math.Abs(got - test.want); diff > 0.01 {
			t.Errorf("got %f != want %f", got, test.want)
		}
	}
}
