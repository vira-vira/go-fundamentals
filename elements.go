package main

import "fmt"

func main() {
	mySlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Printf("Length of slice is %d and capacity is %d\n", len(mySlice), cap(mySlice))
	//for _, i := range mySlice {
	//	fmt.Println(i)
	//}
	fmt.Println(mySlice[4])
	mySlice[2] = 20
	fmt.Println(mySlice[2])
	fmt.Println(mySlice[:3])

}
