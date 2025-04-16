package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, y1, x2, y2 float64
	fmt.Scan(&x1, &y1, &x2, &y2)

	dx := math.Abs(x1 - x2)
	dy := math.Abs(y1 - y2)

	if dx == 0 || dy == 0 || dx == dy {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}
