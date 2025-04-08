package main

import "fmt"

func main() {
	var phonePrice, casePrice, chargerPrice, headphonesPrice uint

	fmt.Scan(&phonePrice, &casePrice, &chargerPrice, &headphonesPrice)

	fullPrice := phonePrice + casePrice + chargerPrice + headphonesPrice
	sumOfBundles := 3

	fmt.Println(fullPrice * uint(sumOfBundles))
}
