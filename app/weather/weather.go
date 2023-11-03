package weather

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

func GetSmallestSpread(file *os.File) int {
	scanner := bufio.NewScanner(file)

	var selectedEntry *WeatherEntry

	for scanner.Scan() {
		row := strings.Join(strings.Fields(scanner.Text()), " ")
		columns := strings.Split(row, " ")

		if entry, err := denormalize(columns); err == nil {
			if selectedEntry == nil {
				selectedEntry = entry

				continue
			}

			if entry.spread < selectedEntry.spread {
				selectedEntry = entry
			}
		}
	}

	return selectedEntry.day
}

func denormalize(data []string) (*WeatherEntry, error) {
	if len(data) < 2 {
		return nil, errors.New("Could not parse column")
	}

	day, err := strconv.Atoi(data[0])

	if err != nil {
		return nil, errors.New("Could not parse day")
	}

	maxTemp, err := strconv.Atoi(data[1])

	if err != nil {
		return nil, errors.New("Could not parse maxTemp")
	}

	minTemp, err := strconv.Atoi(data[2])

	if err != nil {
		return nil, errors.New("Could not parse minTemp")
	}

	return &WeatherEntry{day: day, spread: maxTemp - minTemp}, nil
}

type WeatherEntry struct {
	day    int
	spread int
}
