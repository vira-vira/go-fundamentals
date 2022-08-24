package main

import "fmt"

func main() {
	dddLenghtMins := 275
	cawLengthMins := 30

	if dddLenghtMins > cawLengthMins {
		fmt.Println("Docker deep dive is longer")
	} else if dddLenghtMins == cawLengthMins {
		fmt.Println("time is the same")
	} else {
		fmt.Println("The other is longer")
	}
}
