package factorial

import (
	"fmt"
	"math"
	"math/big"
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
	prec              = 15
	loopLimit float64 = 500000000
)

func (c *Computer) Compute(v int64) string {
	f := float64(v)
	switch {
	case f < loopLimit:
		return factorial(f)
	default:
		return stirlingsApproximation(f)
	}
}
func stirlingsApproximation(n float64) string {
	fmt.Println("stirling")
	p := math.Sqrt(2 * math.Pi * n)
	q := big.NewFloat(math.Pow((n / math.E), n))
	k := big.NewFloat(p)
	return k.Mul(k, q).Text('e', prec)
}

func factorial(x float64) string {
	var prod = big.NewFloat(1)
	for i := float64(0); i < x; i++ {
		prod.Mul(prod, prod.SetFloat64(i))
	}
	return prod.Text('e', prec)
}

// func recursiveFact(x float64) string {

// }

// func main() {

// 	cm := &Computer{}
// 	verybig := cm.Compute(100)

// 	fmt.Println(math.MaxInt64)
// 	fmt.Println(verybig)

// }
