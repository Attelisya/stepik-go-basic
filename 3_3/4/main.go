package main

import "fmt"

func main() {
	var x, y float64
	fmt.Scan(&x, &y)
	if x > 0 {
		if y > 0 {
			fmt.Print(1)
		} else {
			fmt.Print(4)
		}
	} else {
		if x < 0 {
			if y > 0 {
				fmt.Print(2)
			} else {
				fmt.Print(3)
			}
		}
	}
}
