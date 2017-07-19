package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func newton_sqrt(n float64) (float64, int, error) {

	if n < 0 {
		return -1, -1, ErrNegativeSqrt(-2)
	}

	k := float64(1)
	l := 1
	for ; math.Abs(k*k-n) > 1E-3; l++ {
		k = (k + n/k) / 2
	}
	return k, l, nil

}

func main() {

	fmt.Println(newton_sqrt(4))
	fmt.Println(newton_sqrt(-2))
}
