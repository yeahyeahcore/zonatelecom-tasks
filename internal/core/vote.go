package core

type VotingState struct {
	VotingID string            `json:"votingId"`
	Results  []VoteStateResult `json:"results"`
}

type VotingStateOptionsMap struct {
	VotingID string          `json:"votingId"`
	Options  map[string]uint `json:"options"`
}

type VoteStateResult struct {
	OptionID string `json:"optionId"`
	Count    uint   `json:"count"`
}
