package utils_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yeahyeahcore/zonatelecom-tasks/internal/core"
	"github.com/yeahyeahcore/zonatelecom-tasks/internal/utils"
)

func Test_GetVotingPercentageDifference(t *testing.T) {
	testCases := []struct {
		name         string
		inputPrev    *core.VotingStateOptionsMap
		inputCurrent *core.VotingStateOptionsMap
		expected     map[string]float64
	}{
		{
			name: "get voting percentage difference with exuals options id",
			inputPrev: &core.VotingStateOptionsMap{
				VotingID: "1",
				Options: map[string]uint{
					"1": 3,
					"2": 6,
				},
			},
			inputCurrent: &core.VotingStateOptionsMap{
				VotingID: "1",
				Options: map[string]uint{
					"1": 5,
					"2": 15,
				},
			},
			expected: map[string]float64{
				"1": 10,
				"2": 45,
			},
		},
		{
			name: "get voting percentage difference with different options id",
			inputPrev: &core.VotingStateOptionsMap{
				VotingID: "1",
				Options: map[string]uint{
					"1": 3,
					"2": 6,
				},
			},
			inputCurrent: &core.VotingStateOptionsMap{
				VotingID: "1",
				Options: map[string]uint{
					"1": 5,
					"2": 15,
					"3": 2,
				},
			},
			expected: map[string]float64{
				"1": 9,
				"2": 41,
				"3": 9,
			},
		},
		{
			name: "get voting percentage difference with minimum integer change and equals options count",
			inputPrev: &core.VotingStateOptionsMap{
				VotingID: "1",
				Options: map[string]uint{
					"1": 256,
					"2": 267,
				},
			},
			inputCurrent: &core.VotingStateOptionsMap{
				VotingID: "1",
				Options: map[string]uint{
					"1": 257,
					"2": 267,
				},
			},
			expected: map[string]float64{
				"1": 1,
			},
		},
		{
			name: "get voting percentage difference with empty options",
			inputPrev: &core.VotingStateOptionsMap{
				VotingID: "1",
				Options:  map[string]uint{},
			},
			inputCurrent: &core.VotingStateOptionsMap{
				VotingID: "1",
				Options:  map[string]uint{},
			},
			expected: map[string]float64{},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expected, utils.GetVotingPercentageDifference(testCase.inputCurrent, testCase.inputPrev))
		})
	}
}
