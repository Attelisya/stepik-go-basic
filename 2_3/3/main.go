package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	_ = scanner.Scan()
	row1 := scanner.Text()
	_ = scanner.Scan()
	row2 := scanner.Text()
	_ = scanner.Scan()
	row3 := scanner.Text()

	fmt.Print(row1, "\n", row2, "\n", row3)
}
