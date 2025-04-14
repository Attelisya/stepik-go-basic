package main

import "fmt"

func main() {
	var xL, yL, xF, yF int
	fmt.Scan(&xL, &yL, &xF, &yF)
	if xL == xF || yL == yF {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}
