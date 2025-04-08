package main

import "fmt"

func main() {
	const minutesInHour = 60
	var minutes int

	fmt.Scan(&minutes)

	fmt.Printf("%d мин - это %d час %d минут.", minutes, minutes/minutesInHour, minutes%minutesInHour)
}
