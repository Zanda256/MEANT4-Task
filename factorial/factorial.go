package factorial

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
	"strings"
	"sync"
)

type Computer struct {
	prec int
}

func NewComputer(p int) *Computer {
	if p == 0 {
		return &Computer{
			prec: 15,
		}
	}
	return &Computer{prec: p}
}

var (
	prec = 15
	wg   sync.WaitGroup
)

func (c *Computer) Compute(v int64) string {
	f := float64(v)
	switch {
	case f <= 10:
		a := recursiveFact(f)
		return strconv.Itoa(int(a))
	case 10 < f && f <= 170:
		a := recursiveFact(f)
		return strconv.FormatFloat(a, 'g', 15, 64)
	case 170 < f && f <= 1000000:
		return mediumFactorial(v)
	case f > 1000000:
		return veryLargeFactorial(f)
	}
	return ""
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
		go numberOfDecimalDigits(arr[k], arr[k+1], ch1)
	}

	wg.Wait()
	close(ch1)
	var sum float64
	for v := range ch1 {
		sum += v
	}
	totalstr := strconv.FormatFloat((sum), 'g', 15, 64)

	finalStr.WriteString(totalstr)

	return finalStr.String()
}

func numberOfDecimalDigits(x float64, y float64, c chan float64) {
	defer wg.Done()
	var dec float64
	for i := x; i > y; i-- {
		dec += math.Log10(i)
	}
	c <- dec
}

func splitTask(n float64) []float64 {
	gap := n / 10
	arr := make([]float64, 0)
	for i := n; i > -1; i -= gap {
		arr = append(arr, i)
	}
	return arr
}

func mediumFactorial(d int64) string {
	var fact = new(big.Int)
	verybig := fact.MulRange(1, d)
	f := new(big.Float)
	f.SetInt(verybig)
	return f.Text('g', prec)
}

func recursiveFact(x float64) float64 {
	if x == 0 {
		return 1
	}
	return x * recursiveFact(x-1)
}

// func main() {

// 	// cm := &Computer{}
// 	//b := veryLargeFactorial(900000)
// 	//fmt.Println(math.MaxInt64)
// 	// var d float64
// 	// for i := 100; i > 90; i-- {
// 	// 	d += math.Log10(float64(i))
// 	// }
// 	// fmt.Printf("Sum 100-90 = %+v\n", d)
// 	// fmt.Println()
// 	// fmt.Printf("Log 100 = %+v\n", math.Log10(100))
// 	// fmt.Printf("Log 60 = %+v\n", math.Log10(60))
// 	// fmt.Printf("Log 10 = %+v\n", math.Log10(10))
// 	var x float64 = 65000000000
// 	fmt.Println(x)
// 	b := new(big.Float)
// 	fmt.Println(b.SetFloat64(math.MaxInt64 * math.Log10(math.MaxInt64)).String())
// 	st := veryLargeFactorial(10000000000)
// 	fmt.Println(st)
// 	//fmt.Println("log is : ", numberOfDecimalDigits(b))

// 	// fmt.Println(math.MaxInt64)
// 	// fmt.Println(verybig)
// 	// a := recursiveFact(10)
// 	// fmt.Println(int64(a))
// }
