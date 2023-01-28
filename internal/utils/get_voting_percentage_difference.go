package utils

import (
	"math"

	"github.com/yeahyeahcore/zonatelecom-tasks/internal/core"
)

// Returns percentage difference beetwen options counts
func GetVotingPercentageDifference(currentVotingState *core.VotingStateOptionsMap, perviousVotingState *core.VotingStateOptionsMap) map[string]float64 {
	percentageOptionsDifferenceMap := make(map[string]float64)

	totalCount := uint(0)

	for _, count := range currentVotingState.Options {
		totalCount += count
	}

	for currentOptionID, currentOptionCount := range currentVotingState.Options {
		previousPercentage := int(float64(perviousVotingState.Options[currentOptionID]) / float64(totalCount) * 100)
		currentPercentage := int(float64(currentOptionCount) / float64(totalCount) * 100)
		differnce := math.Abs(float64(currentPercentage - previousPercentage))

		if differnce >= 1 {
			percentageOptionsDifferenceMap[currentOptionID] = differnce
		}
	}

	return percentageOptionsDifferenceMap
}
