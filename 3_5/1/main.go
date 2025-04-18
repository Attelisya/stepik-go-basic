package main

import "fmt"

func main() {
	var age int
	var sex string
	fmt.Scan(&age)
	fmt.Scan(&sex)

	if sex == "m" && age >= 12 && age <= 18 {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}
