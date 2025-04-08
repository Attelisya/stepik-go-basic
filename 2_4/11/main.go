package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Printf("Следующее за числом %d число: %d\nДля числа %d предыдущее число: %d", n, n+1, n, n-1)
}
