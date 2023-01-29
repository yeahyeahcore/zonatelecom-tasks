package utils_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/yeahyeahcore/zonatelecom-tasks/internal/core"
	"github.com/yeahyeahcore/zonatelecom-tasks/internal/models"
	"github.com/yeahyeahcore/zonatelecom-tasks/internal/utils"
)

func Test_TransferVotingStatesToCore(t *testing.T) {
	testCases := []struct {
		name     string
		input    []*core.VotingStateOptionsMap
		expected []*core.PreviousVotingState
	}{
		{
			name: "correct transfer voting state model to core with multiple options",
			input: []*core.VotingStateOptionsMap{
				{
					VotingID: "1",
					Options: map[string]uint{
						"1": 2,
						"2": 3,
					},
				},
				{
					VotingID: "2",
					Options: map[string]uint{
						"3": 1,
						"4": 6,
					},
				},
			},
			expected: []*core.PreviousVotingState{
				{
					VotingID: "1",
					Results: []core.VoteStateResult{
						{
							OptionID: "1",
							Count:    2,
						},
						{
							OptionID: "2",
							Count:    3,
						},
					},
				},
				{
					VotingID: "2",
					Results: []core.VoteStateResult{
						{
							OptionID: "3",
							Count:    1,
						},
						{
							OptionID: "4",
							Count:    6,
						},
					},
				},
			},
		},
		{
			name: "correct transfer voting state model to core with single option",
			input: []*core.VotingStateOptionsMap{
				{
					VotingID: "1",
					Options: map[string]uint{
						"1": 2,
					},
				},
			},
			expected: []*core.PreviousVotingState{
				{
					VotingID: "1",
					Results: []core.VoteStateResult{
						{
							OptionID: "1",
							Count:    2,
						},
					},
				},
			},
		},
		{
			name:     "correct transfer voting state model to core with empty array",
			input:    []*core.VotingStateOptionsMap{},
			expected: []*core.PreviousVotingState{},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expected, utils.TransferVotingStatesToCore(testCase.input))
		})
	}
}

func Test_TransferVotingStatesToOptionsMap(t *testing.T) {
	testCases := []struct {
		name     string
		input    []*models.PreviousVotingState
		expected []*core.VotingStateOptionsMap
	}{
		{
			name: "correct transfer voting states to options map with multiple options",
			input: []*models.PreviousVotingState{
				{
					VotingID: "1",
					OptionID: "1",
					Count:    2,
				},
				{
					VotingID: "1",
					OptionID: "2",
					Count:    3,
				},
				{
					VotingID: "2",
					OptionID: "3",
					Count:    1,
				},
				{
					VotingID: "2",
					OptionID: "4",
					Count:    6,
				},
			},
			expected: []*core.VotingStateOptionsMap{
				{
					VotingID: "1",
					Options: map[string]uint{
						"1": 2,
						"2": 3,
					},
				},
				{
					VotingID: "2",
					Options: map[string]uint{
						"3": 1,
						"4": 6,
					},
				},
			},
		},
		{
			name: "correct transfer voting states to options map with single options",
			input: []*models.PreviousVotingState{
				{
					VotingID: "1",
					OptionID: "1",
					Count:    2,
				},
			},
			expected: []*core.VotingStateOptionsMap{
				{
					VotingID: "1",
					Options: map[string]uint{
						"1": 2,
					},
				},
			},
		},
		{
			name:     "correct transfer voting states to options map with empty array",
			input:    []*models.PreviousVotingState{},
			expected: []*core.VotingStateOptionsMap{},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expected, utils.TransferVotingStatesToOptionsMap(testCase.input))
		})
	}
}

func Test_TransferVotingStateToCore(t *testing.T) {
	testCases := []struct {
		name     string
		input    *core.VotingStateOptionsMap
		expected *core.PreviousVotingState
	}{
		{
			name: "correct transfer with difference options",
			input: &core.VotingStateOptionsMap{
				VotingID: "1",
				Options: map[string]uint{
					"1": 2,
					"2": 3,
				},
			},
			expected: &core.PreviousVotingState{
				VotingID: "1",
				Results: []core.VoteStateResult{
					{
						OptionID: "1",
						Count:    2,
					},
					{
						OptionID: "2",
						Count:    3,
					},
				},
			},
		},
		{
			name: "correct transfer voting state with single option",
			input: &core.VotingStateOptionsMap{
				VotingID: "1",
				Options: map[string]uint{
					"1": 2,
				},
			},
			expected: &core.PreviousVotingState{
				VotingID: "1",
				Results: []core.VoteStateResult{
					{
						OptionID: "1",
						Count:    2,
					},
				},
			},
		},
		{
			name:  "correct transfer voting state with empty struct",
			input: &core.VotingStateOptionsMap{},
			expected: &core.PreviousVotingState{
				Results: []core.VoteStateResult{},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expected, utils.TransferVotingStateToCore(testCase.input))
		})
	}
}
