package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	sumN := a + b + c
	prodN := a * b * c
	fmt.Printf("%d %d", sumN, prodN)
}
