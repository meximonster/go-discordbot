CREATE TABLE bets
(
    id            serial PRIMARY KEY,
    team          VARCHAR(100)  NOT NULL,
    prediction    VARCHAR(20)   NOT NULL,
    size          INTEGER  NOT NULL,
    odds          NUMERIC,
    result        VARCHAR(10)   NOT NULL,
    posted_at     TIMESTAMP     DEFAULT (now() at time zone 'utc-3')
);

CREATE TABLE polo_bets
(
    id            serial PRIMARY KEY,
    team          VARCHAR(100)  NOT NULL,
    prediction    VARCHAR(20)   NOT NULL,
    size          INTEGER  NOT NULL,
    odds          NUMERIC,
    result        VARCHAR(10)   NOT NULL,
    posted_at     TIMESTAMP     DEFAULT (now() at time zone 'utc-3')
);

CREATE TABLE users
(
    id            serial PRIMARY KEY,
    alias         VARCHAR(20) NOT NULL,
    discord_id    VARCHAR(50) NOT NULL,
    images        jsonb,
);

CREATE TABLE pets
(
    id            serial PRIMARY KEY,
    alias         VARCHAR(20) NOT NULL,
    images       jsonb,
);

CREATE TABLE emotes
(
    id            serial PRIMARY KEY,
    alias         VARCHAR(20) NOT NULL,
    images        jsonb,
);

CREATE INDEX idx_bets_date ON bets (posted_at);
CREATE INDEX idx_polo_bets_date ON polo_bets (posted_at);
CREATE INDEX idx_users ON users (nickname);
CREATE INDEX idx_pets ON pets (alias);
CREATE INDEX idx_emotes ON emotes (alias);