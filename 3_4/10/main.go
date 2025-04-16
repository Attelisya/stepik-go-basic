package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	var desks1, desks2, desks3 int
	if a%2 == 0 {
		desks1 = a / 2
	} else {
		desks1 = a/2 + 1
	}
	if b%2 == 0 {
		desks2 = b / 2
	} else {
		desks2 = b/2 + 1
	}
	if c%2 == 0 {
		desks3 = c / 2
	} else {
		desks3 = c/2 + 1
	}
	fmt.Print(desks1 + desks2 + desks3)
}
