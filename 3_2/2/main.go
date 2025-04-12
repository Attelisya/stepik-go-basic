package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	if n > -1 && n < 17 {
		fmt.Print("Принадлежит")
	} else {
		fmt.Print("Не принадлежит")
	}
}
