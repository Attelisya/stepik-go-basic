package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	switch n {
	case 12, 1, 2:
		fmt.Print("Зима")
	case 3, 4, 5:
		fmt.Print("Весна")
	case 6, 7, 8:
		fmt.Print("Лето")
	case 9, 10, 11:
		fmt.Print("Осень")
	}
}
