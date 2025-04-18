package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	switch {
	case n <= 3:
		fmt.Print("начинающий")
	case n >= 4 && n <= 7:
		fmt.Print("младший разработчик")
	case n >= 8 && n <= 15:
		fmt.Print("средний разработчик")
	case n >= 16:
		fmt.Print("старший разработчик")
	}
}
