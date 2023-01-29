package core

type CreateVoteRequest struct {
	VoteID   string `json:"voteId"`
	VotingID string `json:"votingId"`
	OptionID string `json:"optionId"`
	Digest   string `json:"digest"`
}

type VotingStateOptionsMap struct {
	VotingID string          `json:"votingId"`
	Options  map[string]uint `json:"options"`
}
