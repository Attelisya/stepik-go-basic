package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	firstNum := n / 100
	secondNum := n % 100 / 10
	thirdNum := n % 10

	fmt.Printf("%d%d%d", thirdNum, secondNum, firstNum)
}
