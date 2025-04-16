package main

import (
	"fmt"
)

func main() {
	var k, m int
	fmt.Scan(&k, &m)

	if m != 0 {
		if k%m == 0 {
			fmt.Print(k / m)
		} else {
			fmt.Print(k/m + 1)
		}
	} else {
		fmt.Print("Вася никогда не решит задачки...")
	}
}
