package parser

import (
	"bufio"
	"os"
	"strings"
)

func FindSmallestDifference(filePath string, selectedKey int, firstKey int, secondKey int) string {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	defer file.Close()

	var selectedEntry *entry

	for scanner.Scan() {
		row := strings.Split(strings.Join(strings.Fields(scanner.Text()), " "), " ")

		currentEntry := denormalize(row, selectedKey, firstKey, secondKey)

		if currentEntry != nil && (selectedEntry == nil || selectedEntry.difference > currentEntry.difference) {
			selectedEntry = currentEntry
		}
	}

	return selectedEntry.title
}
