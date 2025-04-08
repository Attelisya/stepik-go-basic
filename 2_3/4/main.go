package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var lines [3]string

	for i := range 3 {
		if !scanner.Scan() {
			fmt.Fprintln(os.Stderr, "Ошибка чтения ввода")
			return
		}
		lines[i] = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка сканера: %v\n", err)
		return
	}

	fmt.Printf("%s\n%s\n%s\n", lines[2], lines[1], lines[0])
}
