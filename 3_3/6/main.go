/*
Задание со звёздочкой!! :)
Необходимо вычислить квадратный корень,
используя вложенные условия
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64 // a != 0

	fmt.Scan(&a, &b, &c)

	d := math.Pow(b, 2) - 4*a*c

	if d < 0 {
		fmt.Print()
	} else {
		if d == 0 { // считаем один корень
			x := -b / (2 * a)
			fmt.Print(x)
		} else { // считаем 2 корня
			x1 := (-b + math.Sqrt(d)) / (2 * a)
			x2 := (-b - math.Sqrt(d)) / (2 * a)
			/* по условию задачи, если
			уравнение имеет два корня, вывести
			сначала меньший, далее больший */
			if x1 > x2 {
				fmt.Printf("%f\n%f", x2, x1)
			} else {
				fmt.Printf("%f\n%f", x1, x2)
			}
		}
	}
}
