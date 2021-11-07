package factorial

import (
	"fmt"
	"log"
	"math"
	"math/big"
	"os"
	"strconv"
	"strings"
	"sync"
)

type Computer struct{}

func NewComputer() *Computer {
	return &Computer{}
}

var (
	wg   sync.WaitGroup
	prec = setPrecision()
)

func setPrecision() int {
	p, err := strconv.Atoi(os.Getenv("FLOAT_PRECISION"))
	if err != nil {
		log.Printf("can not parse environment variable FLOAT_PRECISION.\n Using default of 15.\n")
		p = 15
	}
	return p
}

func (c *Computer) Compute(v int64) string {
	f := float64(v)
	switch {
	case f <= 10:
		a := recursiveFact(f)
		return strconv.Itoa(int(a))
	case 10 < f && f <= 170:
		a := recursiveFact(f)
		return strconv.FormatFloat(a, 'g', prec, 64)
	case 170 < f && f <= 1000000:
		return mediumFactorial(v)
	case f > 1000000 && f <= 30000000000:
		return veryLargeFactorial(f)
	default:
		return stirlingsApproximation(f)
	}
}

func veryLargeFactorial(x float64) string {
	fmt.Println("working ...\nThis may take a while!")
	finalStr := strings.Builder{}
	finalStr.WriteString("10e+")
	arr := splitTask(x)
	ch1 := make(chan float64, 15)
	for k := range arr {
		if k+1 == 11 {
			break
		}
		wg.Add(1)
		go LogFactorial(arr[k], arr[k+1], ch1)
	}

	wg.Wait()
	close(ch1)
	var sum float64
	for v := range ch1 {
		sum += v
	}
	finalFact := new(big.Float)
	finalFact.SetFloat64(sum)
	finalFact.Mul(finalFact, big.NewFloat(math.Log10(x)))
	totalstr := strconv.FormatFloat((sum), 'g', 15, 64)

	finalStr.WriteString(totalstr)

	return finalStr.String()
}

//Calculate the log of n!
func LogFactorial(x float64, y float64, c chan float64) {
	defer wg.Done()
	var dec float64
	for i := x; i > y; i-- {
		dec += math.Log10(i)
	}
	c <- dec
}

//split the number into 10 intervals to leverage concurrent calculate factorial of each interval
func splitTask(n float64) []float64 {
	gap := n / 10
	arr := make([]float64, 0)
	for i := n; i > -1; i -= gap {
		arr = append(arr, i)
	}
	return arr
}

//Calculate the factorial of a moderately large number
func mediumFactorial(d int64) string {
	var fact = new(big.Int)
	verybig := fact.MulRange(1, d)
	f := new(big.Float)
	f.SetInt(verybig)
	return f.Text('g', prec)
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
	exponentStr := "2.7182818284^"
	lnXFact := ((x * math.Log(x)) - x)
	bLnXFact := big.NewFloat(lnXFact).Text('g', 15)

	return exponentStr + bLnXFact

}
