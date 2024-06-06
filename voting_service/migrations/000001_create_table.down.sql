ALTER TABLE candidates DROP CONSTRAINT candidate_unique_election_public;
ALTER TABLE public_votes DROP CONSTRAINT public_votes_unique_election_public;
DROP table if exists elections, candidates, public_votes, votes cascade;
