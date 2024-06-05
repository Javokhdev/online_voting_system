CREATE TABLE IF NOT EXISTS election(
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    date TIMESTAMP NOT NULL DEFAULT NOW(),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at BIGINT DEFAULT 0
)

CREATE TABLE IF NOT EXISTS candidate(
    id UUID PRIMARY KEY,
    election_id UUID NOT NULL,
    public_id UUID REFERENCES public(id),
    party_id UUID REFERENCES party(id),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at BIGINT DEFAULT 0
)

ALTER TABLE candidate
ADD CONSTRAINT unique_election_public UNIQUE (election_id, public_id); 

CREATE TABLE IF NOT EXISTS public_vote(
    id UUID PRIMARY KEY,
    election_id UUID REFERENCES election(id),
    public_id UUID REFERENCES public(id),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at BIGINT DEFAULT 0
)

ALTER TABLE public_vote
ADD CONSTRAINT unique_election_public UNIQUE (election_id, public_id);

CREATE TABLE IF NOT EXISTS votes(
    id UUID PRIMARY KEY,
    candidate UUID REFERENCES candidate(id),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at BIGINT DEFAULT 0
)