package main

import "fmt"

func main() {
	type courseMeta struct {
		author string
		level  string
		rating float64
	}

	//var meta courseMeta
	//meta := new(courseMeta)

	meta := courseMeta{
		author: "Pepe",
		level:  "intermediate",
		rating: 5,
	}
	fmt.Println(meta)
	fmt.Println("Hola", meta.author)
}
