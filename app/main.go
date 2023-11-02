package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Weather struct {
	dy  string
	mxt string
}

func main() {
	file, err := os.Open("../data/weather.dat")

	if err != nil {
		fmt.Println(err)

		return
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.Join(strings.Fields(scanner.Text()), " ")
		words := strings.Split(line, " ")

		if len(words) != 15 {
			continue
		}

		weather := Weather{dy: words[0], mxt: words[1]}

		fmt.Println(weather)
	}
}
