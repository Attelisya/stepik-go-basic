package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	firD := n / 100000
	secD := n % 100000 / 10000
	thiD := n % 10000 / 1000
	fouD := n % 1000 / 100
	fivD := n % 100 / 10
	sixD := n % 10

	sumF := firD + secD + thiD
	sumL := fouD + fivD + sixD

	if sumF == sumL {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}
