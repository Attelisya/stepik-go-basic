package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	a := x * x // x^2
	b := a * a // x^4
	c := b * a // x^6
	fmt.Println(c)
}
