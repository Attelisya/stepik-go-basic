package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	if a == b {
		fmt.Print(a)
	}
	if a > b {
		fmt.Print(a)
	}
	if b > a {
		fmt.Print(b)
	}
}
