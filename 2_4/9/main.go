package main

import "fmt"

func main() {
	var rubles, kopecks, numberOfPies int

	fmt.Scan(&rubles, &kopecks, &numberOfPies)

	convertedRubles := rubles * 100                              // рубли в копейках
	priceInKopecks := (convertedRubles + kopecks) * numberOfPies // сумма покупки в копейках

	fmt.Printf("%d %d", priceInKopecks/100, priceInKopecks%100)
}
