package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, x2, y1, y2 float64
	fmt.Scan(&x1, &y1, &x2, &y2)
	a := math.Pow(x1-x2, 2)
	b := math.Pow(y1-y2, 2)
	c := a + b
	dotRange := math.Sqrt(c)
	fmt.Print(dotRange)
}
