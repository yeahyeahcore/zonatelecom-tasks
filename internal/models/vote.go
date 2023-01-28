package models

type Vote struct {
	ID       uint   `json:"id" db:"id"`
	VoteID   string `json:"voteId" db:"vote_id"`
	VotingID string `json:"votingId" db:"voting_id"`
	IptionID string `json:"optionId" db:"option_id"`
}
