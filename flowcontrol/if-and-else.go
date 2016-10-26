package main

import (
	"fmt"
	"math"
)

func power(x, n, limit float64) float64 {
	// fmt.Println(x,n)
	// calls to pow are executed and return
	// before the call to fmt.Println in main begins
	if v := math.Pow(x, n); v < limit {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, limit)
	}
	return limit
}

func main() {
	fmt.Println(
		power(3, 2, 20),
		power(3, 3, 20),
	)
}
