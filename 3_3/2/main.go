package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	if n < 60 {
		fmt.Print("Легкий вес")
	} else {
		if n >= 60 {
			if n <= 63 {
				fmt.Print("Первый полусредний вес")
			} else {
				fmt.Print("Полусредний вес")
			}
		}
	}
}
