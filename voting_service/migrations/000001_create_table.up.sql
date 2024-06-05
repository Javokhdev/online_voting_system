CREATE TABLE IF NOT EXISTS elections(
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    date TIMESTAMP NOT NULL DEFAULT NOW(),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at BIGINT DEFAULT 0
);

CREATE TABLE IF NOT EXISTS candidates(
    id UUID PRIMARY KEY,
    election_id UUID REFERENCES elections(id),
    public_id UUID NOT NULL,
    party_id UUID NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at BIGINT DEFAULT 0
);

ALTER TABLE candidates
ADD CONSTRAINT candidate_unique_election_public UNIQUE (election_id, public_id); 

CREATE TABLE IF NOT EXISTS public_votes(
    id UUID PRIMARY KEY,
    election_id UUID REFERENCES elections(id),
    public_id UUID NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at BIGINT DEFAULT 0
);

ALTER TABLE public_votes
ADD CONSTRAINT public_votes_unique_election_public UNIQUE (election_id, public_id);

CREATE TABLE IF NOT EXISTS votes(
    id UUID PRIMARY KEY,
    candidate UUID REFERENCES candidates(id),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at BIGINT DEFAULT 0
);