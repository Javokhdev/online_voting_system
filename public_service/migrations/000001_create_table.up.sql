
CREATE TABLE IF NOT EXISTS parties (
    id UUID PRIMARY KEY,
    name VARCHAR(250) NOT NULL,
    slogan VARCHAR(250) NOT NULL,
    opened_date TIMESTAMP NOT NULL,
    description VARCHAR(250),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at BIGINT DEFAULT 0
);

CREATE TABLE IF NOT EXISTS publics(
    id UUID PRIMARY KEY,
    first_name VARCHAR(250) NOT NULL,
    last_name VARCHAR(250) NOT NULL,
    birthday TIMESTAMP NOT NULL,
    gender VARCHAR(250) NOT NULL,
    nation VARCHAR(250) NOT NULL,
    party_id UUID REFERENCES parties(id),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at BIGINT DEFAULT 0
);