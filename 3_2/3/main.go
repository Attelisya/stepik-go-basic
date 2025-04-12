package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	oneDigit := n / 100
	twoDigit := n % 100 / 10
	threeDigit := n % 10

	if oneDigit != twoDigit && oneDigit != threeDigit && twoDigit != threeDigit {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}
