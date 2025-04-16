package main

import "fmt"

func main() {
	var k2, k3, k5, k6 int
	fmt.Scan(&k2, &k3, &k5, &k6)

	count256 := 0
	count32 := 0

	for k5 > 0 && k2 > 0 && k6 > 0 {
		count256++
		k5--
		k2--
		k6--
	}
	for k3 > 0 && k2 > 0 {
		count32++
		k3--
		k2--
	}
	fmt.Print(256*count256 + 32*count32)
}
