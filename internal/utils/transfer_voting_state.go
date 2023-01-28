package utils

import (
	"github.com/yeahyeahcore/zonatelecom-tasks/internal/core"
	"github.com/yeahyeahcore/zonatelecom-tasks/internal/models"
)

func TransferVotingStateModelsToCore(models []*models.VotingState) []*core.VotingState {
	transferedVotingStateMap := make(map[string][]core.VoteStateResult)

	for _, votingState := range models {
		transferedVotingStateMap[votingState.VotingID] = append(transferedVotingStateMap[votingState.VotingID], core.VoteStateResult{
			OptionID: votingState.OptionID,
			Count:    votingState.Count,
		})
	}

	transferedVotingStateMapIterationNumber := 0
	transferedVotingState := make([]*core.VotingState, len(transferedVotingStateMap))

	for key, result := range transferedVotingStateMap {
		transferedVotingState[transferedVotingStateMapIterationNumber] = &core.VotingState{
			VotingID: key,
			Results:  result,
		}

		transferedVotingStateMapIterationNumber++
	}

	return transferedVotingState
}

func TransferVotingStatesToOptionsMap(votingStates []*core.VotingState) []*core.VotingStateOptionsMap {
	transferedVotingStateOptionsMap := make([]*core.VotingStateOptionsMap, len(votingStates))

	for index, votingState := range votingStates {
		optionsMap := make(map[string]uint)

		for _, count := range votingState.Results {
			optionsMap[count.OptionID] = count.Count
		}

		transferedVotingStateOptionsMap[index] = &core.VotingStateOptionsMap{
			VotingID: votingState.VotingID,
			Options:  optionsMap,
		}
	}

	return transferedVotingStateOptionsMap
}
