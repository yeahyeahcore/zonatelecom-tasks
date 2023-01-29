CREATE TABLE votes (
	id serial PRIMARY KEY,
	vote_id UUID NOT NULL,
	voting_id UUID NOT NULL,
	option_id UUID NOT NULL
);