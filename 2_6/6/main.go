package main

import (
	"fmt"
	"math"
)

func main() {
	var b int
	fmt.Scan(&b)
	kByte := float64(b) / math.Pow(2, 13)
	fmt.Print(kByte)
}
