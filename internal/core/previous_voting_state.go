package core

type PreviousVotingState struct {
	VotingID string            `json:"votingId"`
	Results  []VoteStateResult `json:"results"`
}

type VoteStateResult struct {
	OptionID string `json:"optionId"`
	Count    uint   `json:"count"`
}
