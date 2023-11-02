package main

import "os"

func main() {
	if file, err := os.Open("../data/weather.dat"); err == nil {
		outputSmallestSpread(file)
	}
}
