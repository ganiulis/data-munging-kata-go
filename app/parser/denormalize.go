package parser

import "strconv"

func denormalize(row []string, titleKey int, firstKey int, secondKey int) *entry {
	if len(row) < secondKey {
		return nil
	}

	firstValue, err := strconv.Atoi(row[firstKey])
	if err != nil {
		return nil
	}

	secondValue, err := strconv.Atoi(row[secondKey])
	if err != nil {
		return nil
	}

	var difference int
	if firstValue > secondValue {
		difference = firstValue - secondValue
	} else {
		difference = secondValue - firstValue
	}

	return &entry{title: row[titleKey], difference: difference}
}
