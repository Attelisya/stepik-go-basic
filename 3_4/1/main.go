package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	lastD := n % 10
	if lastD%2 == 0 {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}
