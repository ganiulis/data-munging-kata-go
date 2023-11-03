package main

import (
	"app/football"
	"app/weather"
	"fmt"
	"os"
)

func main() {
	if file, err := os.Open("../data/weather.dat"); err == nil {
		fmt.Printf("Smallest weather spread is on day %d.\n", weather.GetSmallestSpread(file))
	}

	if file, err := os.Open("../data/football.dat"); err == nil {
		fmt.Printf("Smallest football spread is by %s.\n", football.GetSmallestSpreadTeam(file))
	}
}
