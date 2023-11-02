package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("../data/weather.dat")

	if err != nil {
		fmt.Println(err)

		return
	}

	scanner := bufio.NewScanner(file)

	var weathers []Weather

	for scanner.Scan() {
		row := strings.Join(strings.Fields(scanner.Text()), " ")

		if weatherPtr, err := denormalizeWeather(row); err == nil {
			weathers = append(weathers, *weatherPtr)
		}
	}

	fmt.Println(weathers)
}

func denormalizeWeather(row string) (*Weather, error) {
	columns := strings.Split(row, " ")

	if len(columns) < 2 {
		return nil, errors.New("Could not parse column")
	}

	day, err := strconv.Atoi(columns[0])

	if err != nil {
		return nil, errors.New("Could not parse day")
	}

	maxTemp, err := strconv.Atoi(columns[1])

	if err != nil {
		return nil, errors.New("Could not parse maxTemp")
	}

	minTemp, err := strconv.Atoi(columns[2])

	if err != nil {
		return nil, errors.New("Could not parse minTemp")
	}

	return &Weather{day: day, maxTemp: maxTemp, minTemp: minTemp}, nil
}

type Weather struct {
	day     int
	maxTemp int
	minTemp int
}
