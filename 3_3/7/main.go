// Количество дней в месяце

package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	switch n {
	case 1:
		fmt.Print(31)
	case 2:
		fmt.Print(29)
	case 3:
		fmt.Print(31)
	case 4:
		fmt.Print(30)
	case 5:
		fmt.Print(31)
	case 6:
		fmt.Print(30)
	case 7:
		fmt.Print(31)
	case 8:
		fmt.Print(31)
	case 9:
		fmt.Print(30)
	case 10:
		fmt.Print(31)
	case 11:
		fmt.Print(30)
	case 12:
		fmt.Print(31)
	}
}
