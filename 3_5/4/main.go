package main

import (
	"fmt"
	"math"
)

func main() {
	var n float64
	fmt.Scan(&n)

	if n == 0 {
		fmt.Print("Обратного числа не существует")
	} else {
		fmt.Print(math.Pow(n, -1))
	}
}
