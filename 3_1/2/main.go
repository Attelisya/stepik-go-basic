package main

import (
	"fmt"
)

func main() {
	var pass1, pass2 string

	fmt.Scan(&pass1)
	fmt.Scan(&pass2)

	if pass1 == pass2 {
		fmt.Println("Пароль принят")
	} else {
		fmt.Println("Пароль не принят")
	}
}
