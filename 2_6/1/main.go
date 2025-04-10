package main

import "fmt"

func main() {
	var n float32
	fmt.Scan(&n)

	circleArea := 3.14 * n * n

	fmt.Print(circleArea)
}
