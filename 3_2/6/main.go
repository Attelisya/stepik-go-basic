package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	if n <= 13 {
		fmt.Print("детство")
	}
	if n >= 14 && n <= 24 {
		fmt.Print("молодость")
	}
	if n >= 25 && n <= 59 {
		fmt.Print("зрелость")
	}
	if n > 59 {
		fmt.Print("старость")
	}
}
