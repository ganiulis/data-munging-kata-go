package main

import (
	"app/parser"
	"fmt"
	"path/filepath"
)

const WeatherDataPath = "../data/weather.dat"
const FootballDataPath = "../data/football.dat"

func main() {
	if absoluteFilepath, err := filepath.Abs(WeatherDataPath); err == nil {
		fmt.Printf("Day: %s.\n", parser.FindSmallestDifference(absoluteFilepath, 0, 1, 2))
	} else {
		panic(err)
	}

	if absoluteFilepath, err := filepath.Abs(FootballDataPath); err == nil {
		fmt.Printf("Team: %s.\n", parser.FindSmallestDifference(absoluteFilepath, 1, 6, 8))
	} else {
		panic(err)
	}
}
