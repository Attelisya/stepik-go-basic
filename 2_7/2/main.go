package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	a := n % 1000 / 100
	fmt.Print(a)
}
