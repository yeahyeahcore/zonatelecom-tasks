package core

type VoteState struct {
	VotingID string            `json:"votingId"`
	Results  []voteStateResult `json:"results"`
}

type voteStateResult struct {
	OptionID string `json:"optionId"`
	Count    string `json:"count"`
}
