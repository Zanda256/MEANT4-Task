package factorial

import (
	"math"
	"math/big"
	"testing"
)

func TestRecursiveFact(t *testing.T) {
	//Initialization
	testCases := []struct {
		//number whose factorial is to be calculated
		num float64
		//factorial of the number
		expected float64
	}{
		{
			num:      30,
			expected: float64(265252859812191058636308480000000),
		},
		{
			num:      10,
			expected: float64(3628800),
		},
		{
			num:      50,
			expected: float64(30414093201713378043612608166064768844377641568960512000000000000),
		},
		{
			num:      80,
			expected: 7.156945704626 * math.Pow10(118),
		},
		{
			num:      100,
			expected: 9.33262154439 * math.Pow10(157),
		},
		{
			num:      120,
			expected: 6.689502913449 * math.Pow10(198),
		},
		{
			num:      150,
			expected: 5.7133839564458 * math.Pow10(262),
		},
	}

	//execution and validation
	for _, tC := range testCases {
		actual := recursiveFact(tC.num)
		bigActual := new(big.Float).SetPrec(uint(prec)).SetFloat64(actual)

		bigExpected := new(big.Float).SetPrec(uint(prec)).SetFloat64(tC.expected)
		acceptedError := new(big.Float).SetPrec(uint(prec)).SetFloat64(float64(math.Pow10(-prec)))

		diff := bigActual.Sub(bigActual, bigExpected)
		diffAbs := diff.Abs(diff)

		if less := diffAbs.Cmp(acceptedError); less > 0 {
			t.Errorf("expected factorial of %+v to be %+v, got %+v", tC.num, tC.expected, actual)
		}
	}

}

func TestMediumFactorial(t *testing.T) {
	testCases := []struct {
		//number whose factorial is to be calculated
		num int64
		//factorial of the number
		expected string
	}{
		{
			num:      1000,
			expected: "4.02387260077093773543702433923e+2567",
		},
		{
			num:      10000,
			expected: "2.846259680917054519e+35659",
		},
		{
			num:      1000000,
			expected: "8.26393168833124006237664610317e+5565708",
		},
	}
	for _, tC := range testCases {
		actual := mediumFactorial(tC.num)
		bigExpected, ok := new(big.Float).SetPrec(uint(prec)).SetString(tC.expected)
		if !ok {
			t.Errorf("could not parse expected factorial string %+v: %+v", tC.num, tC.expected)
		}
		bigActual, ok := new(big.Float).SetPrec(uint(prec)).SetString(actual)
		if !ok {
			t.Errorf("could not parse actual factorial string %+v: %+v", tC.num, actual)
		}
		diff := bigActual.Sub(bigActual, bigExpected)
		diffAbs := diff.Abs(diff)

		acceptedError := new(big.Float).SetPrec(uint(prec)).SetFloat64(float64(math.Pow10(-prec)))
		if less := diffAbs.Cmp(acceptedError); less > 0 {
			t.Errorf("expected factorial of %+v to be %+v, got %+v", tC.num, tC.expected, actual)
		}
	}

}

func TestStirlingsApproximation(t *testing.T) {
	testCases := []struct {
		//number whose factorial is to be calculated
		num int64
		//factorial of the number
		expected float64
	}{
		{
			num:      7000000000,
			expected: ((float64(7000000000) * math.Log(float64(7000000000))) - float64(7000000000)),
		},
		{
			num:      10000,
			expected: ((float64(10000) * math.Log(float64(10000))) - float64(10000)),
		},
		{
			num:      1000000,
			expected: ((float64(1000000) * math.Log(float64(1000000))) - float64(1000000)),
		},
	}
	for _, tC := range testCases {
		actual := stirlingsApproximation(float64(tC.num))

		bigExpected := new(big.Float).SetPrec(uint(prec)).SetFloat64(tC.expected)

		bigActual, ok := new(big.Float).SetPrec(uint(prec)).SetString(actual)
		if !ok {
			t.Errorf("could not parse actual Nln(N) - N  string %+v: %+v", tC.num, actual)
		}

		diff := bigActual.Sub(bigActual, bigExpected)
		diffAbs := diff.Abs(diff)

		acceptedError := new(big.Float).SetPrec(uint(prec)).SetFloat64(float64(math.Pow10(-prec)))
		if less := diffAbs.Cmp(acceptedError); less > 0 {
			t.Errorf("expected Nln(N) - N of %+v to be %+v, got %+v", tC.num, tC.expected, actual)
		}
	}

}

//1000! = 4.02387260077093773543702433923e+2567
//10000! = +2.846259680917054519e+35659

//7000000000! = 1.25e+65875624912
