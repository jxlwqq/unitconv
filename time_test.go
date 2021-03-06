package unitconv

import (
	"math"
	"testing"
)

func TestTime(t *testing.T) {
	tests := []struct {
		fromValue float64
		fromUnit  int
		toUnit    int
		want      float64
	}{
		{1, Second, Second, 1},
		{1, Second, Millisecond, 1000},
		{1, Second, Microsecond, 1000000},
		{1, Second, Nanosecond, 1000000000},
		{1, Second, Picosecond, 1000000000000},
		{1, Second, Minute, 0.0166666667},
		{1, Second, Hour, 0.0002777778},
		{1, Second, Day, 0.0000115741},
		{1, Second, Week, 0.0000016534},
		{1, Second, Month, 3.802570537e-7},
		{1, Second, Year, 3.168808781e-8},
		{1, Millisecond, Second, 0.001},
		{1, Millisecond, Millisecond, 1},
		{1, Millisecond, Microsecond, 1000},
		{1, Millisecond, Nanosecond, 1000000},
		{1, Millisecond, Picosecond, 1000000000},
		{1, Millisecond, Minute, 0.0000166667},
		{1, Millisecond, Hour, 2.777777777e-7},
		{1, Millisecond, Day, 1.157407407e-8},
		{1, Millisecond, Week, 1.653439153e-9},
		{1, Millisecond, Month, 3.802570537e-10},
		{1, Millisecond, Year, 3.168808781e-11},
		{1, Microsecond, Second, 0.000001},
		{1, Microsecond, Millisecond, 0.001},
		{1, Microsecond, Microsecond, 1},
		{1, Microsecond, Nanosecond, 1000},
		{1, Microsecond, Picosecond, 1000000},
		{1, Microsecond, Minute, 1.666666666e-8},
		{1, Microsecond, Hour, 2.777777777e-10},
		{1, Microsecond, Day, 1.157407407e-11},
		{1, Microsecond, Week, 1.653439153e-12},
		{1, Microsecond, Month, 3.802570537e-13},
		{1, Microsecond, Year, 3.168808781e-14},
		{1, Nanosecond, Second, 1.e-9},
		{1, Nanosecond, Millisecond, 0.000001},
		{1, Nanosecond, Microsecond, 0.001},
		{1, Nanosecond, Nanosecond, 1},
		{1, Nanosecond, Picosecond, 1000},
		{1, Nanosecond, Minute, 1.666666666e-11},
		{1, Nanosecond, Hour, 2.777777777e-13},
		{1, Nanosecond, Day, 1.157407407e-14},
		{1, Nanosecond, Week, 1.653439153e-15},
		{1, Nanosecond, Month, 3.802570537e-16},
		{1, Nanosecond, Year, 3.168808781e-17},
		{1, Picosecond, Second, 1.e-12},
		{1, Picosecond, Millisecond, 1.e-9},
		{1, Picosecond, Microsecond, 0.000001},
		{1, Picosecond, Nanosecond, 0.001},
		{1, Picosecond, Picosecond, 1},
		{1, Picosecond, Minute, 1.666666666e-14},
		{1, Picosecond, Hour, 2.777777777e-16},
		{1, Picosecond, Day, 1.157407407e-17},
		{1, Picosecond, Week, 1.653439153e-18},
		{1, Picosecond, Month, 3.802570537e-19},
		{1, Picosecond, Year, 3.168808781e-20},
		{1, Minute, Second, 60},
		{1, Minute, Millisecond, 60000},
		{1, Minute, Microsecond, 60000000},
		{1, Minute, Nanosecond, 60000000000},
		{1, Minute, Picosecond, 60000000000000},
		{1, Minute, Minute, 1},
		{1, Minute, Hour, 0.0166666667},
		{1, Minute, Day, 0.0006944444},
		{1, Minute, Week, 0.0000992063},
		{1, Minute, Month, 0.0000228154},
		{1, Minute, Year, 0.0000019013},
		{1, Hour, Second, 3600},
		{1, Hour, Millisecond, 3600000},
		{1, Hour, Microsecond, 3600000000},
		{1, Hour, Nanosecond, 3600000000000},
		{1, Hour, Picosecond, 3600000000000000},
		{1, Hour, Minute, 60},
		{1, Hour, Hour, 1},
		{1, Hour, Day, 0.0416666667},
		{1, Hour, Week, 0.005952381},
		{1, Hour, Month, 0.0013689254},
		{1, Hour, Year, 0.0001140771},
		{1, Day, Second, 86400},
		{1, Day, Millisecond, 86400000},
		{1, Day, Microsecond, 86400000000},
		{1, Day, Nanosecond, 86400000000000},
		{1, Day, Picosecond, 86400000000000000},
		{1, Day, Minute, 1440},
		{1, Day, Hour, 24},
		{1, Day, Day, 1},
		{1, Day, Week, 0.1428571429},
		{1, Day, Month, 0.0328542094},
		{1, Day, Year, 0.0027378508},
		{1, Week, Second, 604800},
		{1, Week, Millisecond, 604800000},
		{1, Week, Microsecond, 604800000000},
		{1, Week, Nanosecond, 604800000000000},
		{1, Week, Picosecond, 604800000000000000},
		{1, Week, Minute, 10080},
		{1, Week, Hour, 168},
		{1, Week, Day, 7},
		{1, Week, Week, 1},
		{1, Week, Month, 0.2299794661},
		{1, Week, Year, 0.0191649555},
		{1, Month, Second, 2629800},
		{1, Month, Millisecond, 2629800000},
		{1, Month, Microsecond, 2629800000000},
		{1, Month, Nanosecond, 2629800000000000},
		{1, Month, Picosecond, 2629800000000000000},
		{1, Month, Minute, 43830},
		{1, Month, Hour, 730.5},
		{1, Month, Day, 30.4375},
		{1, Month, Week, 4.3482142857},
		{1, Month, Month, 1},
		{1, Month, Year, 0.0833333333},
		{1, Year, Second, 31557600},
		{1, Year, Millisecond, 31557600000},
		{1, Year, Microsecond, 31557600000000},
		{1, Year, Nanosecond, 31557600000000000},
		{1, Year, Picosecond, 31557600000000000000},
		{1, Year, Minute, 525960},
		{1, Year, Hour, 8766},
		{1, Year, Day, 365.25},
		{1, Year, Week, 52.178571429},
		{1, Year, Month, 12},
		{1, Year, Year, 1},
	}

	for _, test := range tests {
		conv := NewTime(test.fromValue, test.fromUnit)
		got, _ := conv.GetToValue(test.toUnit)
		if diff := math.Abs(got - test.want); diff > 0.01 {
			t.Errorf("got %f != want %f", got, test.want)
		}
	}
}
