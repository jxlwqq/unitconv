package unitconv

import (
	"math"
	"testing"
)

func TestLength(t *testing.T) {
	tests := []struct {
		fromValue float64
		fromUnit  int
		toUnit    int
		want      float64
	}{
		{1, Meter, Meter, 1},
		{1, Meter, Kilometer, 0.001},
		{1, Meter, Centimeter, 100},
		{1, Meter, Millimeter, 1000},
		{1, Meter, Micrometer, 1000000},
		{1, Meter, Nanometer, 1000000000},
		{1, Meter, Mile, 0.0006213689},
		{1, Meter, Yard, 1.0936132983},
		{1, Meter, Foot, 3.280839895},
		{1, Meter, Inch, 39.37007874},
		{1, Meter, LightYear, 1.057008707e-16},
		{1, Kilometer, Meter, 1000},
		{1, Kilometer, Kilometer, 1},
		{1, Kilometer, Centimeter, 100000},
		{1, Kilometer, Millimeter, 1000000},
		{1, Kilometer, Micrometer, 1000000000},
		{1, Kilometer, Nanometer, 1000000000000},
		{1, Kilometer, Mile, 0.6213688756},
		{1, Kilometer, Yard, 1093.6132983},
		{1, Kilometer, Foot, 3280.839895},
		{1, Kilometer, Inch, 39370.07874},
		{1, Kilometer, LightYear, 1.057008707e-13},
		{1, Centimeter, Meter, 0.01},
		{1, Centimeter, Kilometer, 0.00001},
		{1, Centimeter, Centimeter, 1},
		{1, Centimeter, Millimeter, 10},
		{1, Centimeter, Micrometer, 10000},
		{1, Centimeter, Nanometer, 10000000},
		{1, Centimeter, Mile, 0.0000062137},
		{1, Centimeter, Yard, 0.010936133},
		{1, Centimeter, Foot, 0.032808399},
		{1, Centimeter, Inch, 0.3937007874},
		{1, Centimeter, LightYear, 1.057008707e-18},
		{1, Millimeter, Meter, 0.001},
		{1, Millimeter, Kilometer, 0.000001},
		{1, Millimeter, Centimeter, 0.1},
		{1, Millimeter, Millimeter, 1},
		{1, Millimeter, Micrometer, 1000},
		{1, Millimeter, Nanometer, 1000000},
		{1, Millimeter, Mile, 6.213688756e-7},
		{1, Millimeter, Yard, 0.0010936133},
		{1, Millimeter, Foot, 0.0032808399},
		{1, Millimeter, Inch, 0.0393700787},
		{1, Millimeter, LightYear, 1.057008707e-19},
		{1, Micrometer, Meter, 0.000001},
		{1, Micrometer, Kilometer, 1.e-9},
		{1, Micrometer, Centimeter, 0.0001},
		{1, Micrometer, Millimeter, 0.001},
		{1, Micrometer, Micrometer, 1},
		{1, Micrometer, Nanometer, 1000},
		{1, Micrometer, Mile, 6.213688756e-10},
		{1, Micrometer, Yard, 0.0000010936},
		{1, Micrometer, Foot, 0.0000032808},
		{1, Micrometer, Inch, 0.0000393701},
		{1, Micrometer, LightYear, 1.057008707e-22},
		{1, Nanometer, Meter, 1.e-9},
		{1, Nanometer, Kilometer, 1.e-12},
		{1, Nanometer, Centimeter, 1.e-7},
		{1, Nanometer, Millimeter, 0.000001},
		{1, Nanometer, Micrometer, 0.001},
		{1, Nanometer, Nanometer, 1},
		{1, Nanometer, Mile, 6.213688756e-13},
		{1, Nanometer, Yard, 1.093613298e-9},
		{1, Nanometer, Foot, 3.280839895e-9},
		{1, Nanometer, Inch, 3.937007874e-8},
		{1, Nanometer, LightYear, 1.057008707e-25},
		{1, Mile, Meter, 1609.35},
		{1, Mile, Kilometer, 1.60935},
		{1, Mile, Centimeter, 160935},
		{1, Mile, Millimeter, 1609350},
		{1, Mile, Micrometer, 1609350000},
		{1, Mile, Nanometer, 1609350000000},
		{1, Mile, Mile, 1},
		{1, Mile, Yard, 1760.0065617},
		{1, Mile, Foot, 5280.019685},
		{1, Mile, Inch, 63360.23622},
		{1, Mile, LightYear, 1.701096963e-13},
		{1, Yard, Meter, 0.9144},
		{1, Yard, Kilometer, 0.0009144},
		{1, Yard, Centimeter, 91.44},
		{1, Yard, Millimeter, 914.4},
		{1, Yard, Micrometer, 914400},
		{1, Yard, Nanometer, 914400000},
		{1, Yard, Mile, 0.0005681797},
		{1, Yard, Yard, 1},
		{1, Yard, Foot, 3},
		{1, Yard, Inch, 36},
		{1, Yard, LightYear, 9.665287622e-17},
		{1, Foot, Meter, 0.3048},
		{1, Foot, Kilometer, 0.0003048},
		{1, Foot, Centimeter, 30.48},
		{1, Foot, Millimeter, 304.8},
		{1, Foot, Micrometer, 304800},
		{1, Foot, Nanometer, 304800000},
		{1, Foot, Mile, 0.0001893932},
		{1, Foot, Yard, 0.3333333333},
		{1, Foot, Foot, 1},
		{1, Foot, Inch, 12},
		{1, Foot, LightYear, 3.22176254e-17},
		{1, Inch, Meter, 0.0254},
		{1, Inch, Kilometer, 0.0000254},
		{1, Inch, Centimeter, 2.54},
		{1, Inch, Millimeter, 25.4},
		{1, Inch, Micrometer, 25400},
		{1, Inch, Nanometer, 25400000},
		{1, Inch, Mile, 0.0000157828},
		{1, Inch, Yard, 0.0277777778},
		{1, Inch, Foot, 0.0833333333},
		{1, Inch, Inch, 1},
		{1, Inch, LightYear, 2.684802117e-18},
		{1, LightYear, Meter, 9460660000000000},
		{1, LightYear, Kilometer, 9460660000000},
		{1, LightYear, Centimeter, 946066000000000000},
		{1, LightYear, Millimeter, 9460660000000000000},
		{1, LightYear, Micrometer, 9.46066e+21},
		{1, LightYear, Nanometer, 9.46066e+24},
		{1, LightYear, Mile, 5878559666946.283203},
		{1, LightYear, Yard, 10346303587051618},
		{1, LightYear, Foot, 31038910761154856},
		{1, LightYear, Inch, 372466929133858300},
		{1, LightYear, LightYear, 1},
	}

	for _, test := range tests {
		conv := NewLength(test.fromValue, test.fromUnit)
		got, _ := conv.GetToValue(test.toUnit)
		if diff := math.Abs(got - test.want); diff > 0.0001 {
			t.Errorf("got %f != want %f", got, test.want)
		}
	}
}
