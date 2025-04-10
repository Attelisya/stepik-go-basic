package main

import "fmt"

func main() {
	var a, b float64
	fmt.Scan(&a, &b)
	s := (a * b) / 2
	fmt.Print(s)
}
