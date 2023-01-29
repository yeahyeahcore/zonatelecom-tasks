package utils

import (
	"github.com/yeahyeahcore/zonatelecom-tasks/internal/core"
	"github.com/yeahyeahcore/zonatelecom-tasks/internal/models"
)

func TransferVotingStatesToCore(votingStates []*core.VotingStateOptionsMap) []*core.PreviousVotingState {
	transferedVotingState := make([]*core.PreviousVotingState, len(votingStates))

	for index, votingState := range votingStates {
		transferedVotingState[index] = TransferVotingStateToCore(votingState)
	}

	return transferedVotingState
}

func TransferVotingStateToCore(votingState *core.VotingStateOptionsMap) *core.PreviousVotingState {
	results := make([]core.VoteStateResult, len(votingState.Options))
	optionsIterationNumber := 0

	for optionID, optionCount := range votingState.Options {
		results[optionsIterationNumber] = core.VoteStateResult{
			OptionID: optionID,
			Count:    optionCount,
		}

		optionsIterationNumber++
	}

	return &core.PreviousVotingState{
		VotingID: votingState.VotingID,
		Results:  results,
	}
}

func TransferVotingStatesToOptionsMap(votingStates []*models.PreviousVotingState) []*core.VotingStateOptionsMap {
	transferedVotingStateMap := make(map[string][]core.VoteStateResult)

	for _, votingState := range votingStates {
		transferedVotingStateMap[votingState.VotingID] = append(transferedVotingStateMap[votingState.VotingID], core.VoteStateResult{
			OptionID: votingState.OptionID,
			Count:    votingState.Count,
		})
	}

	transferedVotingStateMapIterationNumber := 0
	transferedVotingStateOptionsMap := make([]*core.VotingStateOptionsMap, len(transferedVotingStateMap))

	for key, results := range transferedVotingStateMap {
		optionsMap := make(map[string]uint)

		for _, result := range results {
			optionsMap[result.OptionID] += result.Count
		}

		transferedVotingStateOptionsMap[transferedVotingStateMapIterationNumber] = &core.VotingStateOptionsMap{
			VotingID: key,
			Options:  optionsMap,
		}

		transferedVotingStateMapIterationNumber++
	}

	return transferedVotingStateOptionsMap
}
