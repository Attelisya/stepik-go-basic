package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	sign(x)
}

func sign(x int) {
	if x < 0 {
		fmt.Print(-1)
	}
	if x == 0 {
		fmt.Print(0)
	}
	if x > 0 {
		fmt.Print(1)
	}
}
