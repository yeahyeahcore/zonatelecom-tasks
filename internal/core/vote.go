package core

type CreateVoteRequest struct {
	VoteID   string `json:"voteId"`
	VotingID string `json:"votingId"`
	OptionID string `json:"optionId"`
	Digest   string `json:"digest"`
}

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
