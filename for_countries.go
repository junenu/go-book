package main

import "fmt"

func main() {
	countries := [...]string{"Japan", "USA", "China"}

	for _, country := range countries {
		fmt.Println(country)
	}
}
