package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	a := n % 1000
	firstNum := a / 100
	secondNum := a % 100 / 10
	thirdNum := a % 10
	sumNum := firstNum + secondNum + thirdNum
	fmt.Print(sumNum)
}
