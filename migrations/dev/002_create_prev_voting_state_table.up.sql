CREATE TABLE previous_voting_states (
	id serial PRIMARY KEY,
	voting_id UUID NOT NULL,
	option_id UUID NOT NULL,
  count INT NOT NULL
);