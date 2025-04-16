package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	nextN := n + 1
	if nextN%2 != 0 {
		nextN++
	}
	fmt.Print(nextN)
}
