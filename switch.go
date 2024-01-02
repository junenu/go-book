package main

import "fmt"

func main() {
	for i:=0;i<5;i++{
		switch i {
		case 0:
			fmt.Println("This is 0.")
		case 1,4:
			fmt.Println("1 or 4")
		default:
			fmt.Println("yeah")
		}
	}
}