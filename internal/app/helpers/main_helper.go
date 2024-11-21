package helpers

import (
	"fmt"
	"time"
)

func ValidateDate(input string) (time.Time, error) {
	const dateFormat = "02-01-2006"

	parsedTime, err := time.Parse(dateFormat, input)
	if err != nil {

		return time.Time{}, fmt.Errorf("invalid date format: %s", input)
	}
	return parsedTime, nil
}
