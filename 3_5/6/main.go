package main

import "fmt"

func main() {
	var d1, d2, d3 int
	fmt.Scan(&d1, &d2, &d3)

	minDist := min(d1+d2+d3, 2*(d1+d2), 2*(d1+d3), 2*(d2+d3))

	fmt.Println(minDist)
}
