package main

import "fmt"

func main() {
	frases := []string{
		"Pepe",
		"honguito",
		"comio",
		"un mosquito",
		"hola",
	}
	masfrases := []string{
		"hola",
		"manola",
	}
	for _, i := range frases {
		fmt.Println(i)
		for _, j := range masfrases {
			if i == j {
				fmt.Println("esto es mas frases", j)
			}
		}
	}
}
