package main

import "fmt"

func main() {
	var a, b int
	var operator string

	fmt.Scan(&a, &b)
	fmt.Scan(&operator)

	if operator == "+" {
		fmt.Print(a + b)
	} else {
		if operator == "-" {
			fmt.Print(a - b)
		} else {
			if operator == "*" {
				fmt.Print(a * b)
			} else {
				if operator == "/" {
					if b != 0 {
						fmt.Print(float64(a) / float64(b))
					} else {
						fmt.Print("На ноль делить нельзя!")
					}
				} else {
					fmt.Print("Неверная операция")
				}
			}
		}
	}
}
