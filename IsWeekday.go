package main
import (
	"fmt"
	"time"
)

const (
	Sunday    = 0
	Monday    = 1
	Tuesday   = 2
	Wednesday = 3
	Thursday  = 4
	Friday    = 5
	Saturday  = 6
)

func main() {
	d := time.Now().Weekday()
	if d == 0 || d == 6 {
		fmt.Println("週末です")
	} else {
		fmt.Println("平日です！！！")
	}
}