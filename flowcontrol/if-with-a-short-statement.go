package main

import (
	"fmt"
	"math"
)

func power(x, n, limit float64) float64 {
	if v := math.Pow(x, n); v < limit {
		return v
	}
	return limit
}

func main() {
	fmt.Println(power(3, 2, 20))
	fmt.Println(power(3, 3, 20))
}
