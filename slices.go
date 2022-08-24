package main

import "fmt"

func main() {
	//courses := make([]string, 5, 10)
	courses := []string{"uno", "dos", "tres"}

	fmt.Printf("Length of slice is %d and capacity is %d\n", len(courses), cap(courses))
	for _, i := range courses {
		fmt.Println(i)
	}

}
