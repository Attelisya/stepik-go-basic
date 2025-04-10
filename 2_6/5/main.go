package main

import "fmt"

func main() {
	var x float64
	fmt.Scan(&x)
	fract := x - float64(int(x))
	fmt.Print(fract)
}
