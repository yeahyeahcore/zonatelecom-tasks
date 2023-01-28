package utils

import (
	"math"
)

// Returns percentage difference beetwen options counts
func GetVotingPercentageDifference(currentOptions map[string]uint, perviousOptions map[string]uint) map[string]float64 {
	percentageOptionsDifferenceMap := make(map[string]float64)

	totalCount := uint(0)

	for _, count := range currentOptions {
		totalCount += count
	}

	for currentOptionID, currentOptionCount := range currentOptions {
		previousPercentage := int(float64(perviousOptions[currentOptionID]) / float64(totalCount) * 100)
		currentPercentage := int(float64(currentOptionCount) / float64(totalCount) * 100)
		differnce := math.Abs(float64(currentPercentage - previousPercentage))

		if differnce >= 1 {
			percentageOptionsDifferenceMap[currentOptionID] = differnce
		}
	}

	return percentageOptionsDifferenceMap
}
