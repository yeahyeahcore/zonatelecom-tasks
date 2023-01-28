package parse

import (
	"errors"
	"fmt"
	"regexp"
	"time"
)

// Parsing STD UTC value as +03:00
func UTCString(duration string) (time.Time, error) {
	regexSTDWithoutPlusMinus, err := regexp.Compile("^(0[0-9]|1[0-9]|2[0-3]):(0[0-9]|[1-5][0-9])$")
	if err != nil {
		return time.Now().In(time.UTC), fmt.Errorf("failed to create regex to check std offset without plus or minus format %s", err.Error())
	}

	regexSTD, err := regexp.Compile("^[+-](0[0-9]|1[0-9]|2[0-3]):(0[0-9]|[1-5][0-9])$")
	if err != nil {
		return time.Now().In(time.UTC), fmt.Errorf("failed to create regex to check std offset format %s", err.Error())
	}

	isSTDFormat := regexSTD.MatchString(duration)
	if !isSTDFormat {
		isSTDFormatWithoutPlusMinus := regexSTDWithoutPlusMinus.MatchString(duration)

		if isSTDFormatWithoutPlusMinus {
			return time.Now().In(time.UTC), nil
		}

		return time.Now().In(time.UTC), errors.New("invalid std offset format")
	}

	durationMatches := regexSTD.FindStringSubmatch(duration)

	parsedDuration, err := time.ParseDuration(fmt.Sprintf("%sh%sm", durationMatches[1], durationMatches[2]))
	if err != nil {
		return time.Now().In(time.UTC), fmt.Errorf("failed to parse time duration: %s", err.Error())
	}

	return time.Now().In(time.UTC).Add(parsedDuration), nil
}
