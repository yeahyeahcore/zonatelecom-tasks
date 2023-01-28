package models

type Vote struct {
	ID       uint   `json:"id" db:"id"`
	VoteID   string `json:"voteId" db:"vote_id"`
	VotingID string `json:"votingId" db:"voting_id"`
	OptionID string `json:"optionId" db:"option_id"`
}

type VotingState struct {
	VotingID string `json:"votingId" db:"voting_id"`
	OptionID string `json:"optionId" db:"option_id"`
	Count    uint   `json:"count" db:"option_count"`
}
