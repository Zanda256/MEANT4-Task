package factorial

import (
	"log"
	"math"
	"math/big"
	"os"
	"strconv"
)

type Computer struct{}

func NewComputer() *Computer {
	return &Computer{}
}

var (
	prec = setPrecision()
)

func setPrecision() int {
	p, err := strconv.Atoi(os.Getenv("FLOAT_PRECISION"))
	if err != nil {
		log.Printf("can not parse environment variable FLOAT_PRECISION.\n Using default of 30.\n")
		p = 30
	}
	return p
}

func (c *Computer) Compute(v int64) string {
	f := float64(v)
	switch {

	case f <= 170:
		a := recursiveFact(f)
		return strconv.FormatFloat(a, 'g', prec, 64)

	case 170 < f && f <= 1000000:
		fltFact := mediumFactorial(v)
		return fltFact

	default:
		exponentStr := "2.7182818284^"
		lnNfact := stirlingsApproximation(f)
		return exponentStr + lnNfact
	}
}

//Calculate the factorial of a moderately large number
func mediumFactorial(d int64) string {
	//use a big.Int to calculate factorial of d (MulRange)
	var fact = new(big.Int)
	verybig := fact.MulRange(1, d)
	f := new(big.Float)
	//convert the factorial into a big.Float for easy readability
	flt := f.SetInt(verybig)
	return flt.Text('g', prec)
}

//use recursion to calculate factorial of small numbers.
func recursiveFact(x float64) float64 {
	if x == 0 {
		return 1
	}
	return x * recursiveFact(x-1)
}

func stirlingsApproximation(x float64) string {
	//stirling approx : lnN! = N*lnN - N
	lnXFact := ((x * math.Log(x)) - x)
	bigLnXFact := big.NewFloat(lnXFact).Text('g', prec)

	return bigLnXFact

}
