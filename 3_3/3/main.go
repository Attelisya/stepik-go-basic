package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	count := 0
	if a == b {
		count = 2
		if a == c {
			count = 3
		}
	} else {
		if b == c {
			count = 2
		} else {
			if a == c {
				count = 2
			}
		}
	}
	fmt.Print(count)
}
