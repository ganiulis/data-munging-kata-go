package football

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

func GetSmallestSpreadTeam(file *os.File) string {
	scanner := bufio.NewScanner(file)

	var selectedEntry *TeamScore

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

	return selectedEntry.name
}

func denormalize(data []string) (*TeamScore, error) {
	if len(data) < 10 {
		return nil, errors.New("Could not parse column")
	}

	forScore, err := strconv.Atoi(data[6])

	if err != nil {
		return nil, errors.New("Could not parse forScore")
	}

	againstScore, err := strconv.Atoi(data[8])

	if err != nil {
		return nil, errors.New("Could not parse againstScore")
	}

	var spread int

	if forScore > againstScore {
		spread = forScore - againstScore
	} else {
		spread = againstScore - forScore
	}

	return &TeamScore{name: data[1], spread: spread}, nil
}

type TeamScore struct {
	name   string
	spread int
}
