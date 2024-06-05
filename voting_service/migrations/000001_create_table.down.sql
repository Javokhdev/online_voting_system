ALTER TABLE candidate DROP CONSTRAINT unique_election_public;
ALTER TABLE public_vote DROP CONSTRAINT unique_election_public;
DROP table if exists election, candidate, public_vote, votes cascade;
