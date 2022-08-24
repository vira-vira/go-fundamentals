package main

import "fmt"

func main() {
	leaguesTitles := make(map[string]int)
	leaguesTitles["River"] = 10
	leaguesTitles["Boca"] = 8
	recentMatchesWins := map[string]int{
		"River": 2,
		"Boca":  0,
	}
	fmt.Printf("Titulos: %v \nPartidos recientes: %v \n", leaguesTitles, recentMatchesWins)

	for mapKey, mapVal := range leaguesTitles {
		fmt.Printf("key is: %v, Value is: %v\n", mapKey, mapVal)
	}
}
