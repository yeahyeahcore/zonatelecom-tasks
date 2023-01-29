package models

type PreviousVotingState struct {
	ID       uint   `json:"id" db:"id"`
	VotingID string `json:"votingId" db:"voting_id"`
	OptionID string `json:"optionId" db:"option_id"`
	Count    uint   `json:"count" db:"count"`
}
