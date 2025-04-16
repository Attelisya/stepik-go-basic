package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	firD := n / 1000
	secD := n % 1000 / 100
	thiD := n % 100 / 10
	fouD := n % 10

	if firD == fouD && secD == thiD {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}
