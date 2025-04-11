package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	if a%b == 0 {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}
