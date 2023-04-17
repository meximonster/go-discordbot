CREATE TABLE bets
(
    id            serial PRIMARY KEY,
    team          VARCHAR(100)  NOT NULL,
    prediction    VARCHAR(20)   NOT NULL,
    size          INTEGER  NOT NULL,
    odds          NUMERIC,
    result        VARCHAR(10)   NOT NULL,
    posted_at     TIMESTAMP     DEFAULT now()
);

CREATE TABLE polo_bets
(
    id            serial PRIMARY KEY,
    team          VARCHAR(100)  NOT NULL,
    prediction    VARCHAR(20)   NOT NULL,
    size          INTEGER  NOT NULL,
    odds          NUMERIC,
    result        VARCHAR(10)   NOT NULL,
    posted_at     TIMESTAMP     DEFAULT now()
);

CREATE TABLE nick_bets
(
    id            serial PRIMARY KEY,
    team          VARCHAR(100)  NOT NULL,
    prediction    VARCHAR(20)   NOT NULL,
    size          INTEGER  NOT NULL,
    odds          NUMERIC,
    result        VARCHAR(10)   NOT NULL,
    posted_at     TIMESTAMP     DEFAULT now()
);

CREATE TABLE panos_bets
(
    id            serial PRIMARY KEY,
    team          VARCHAR(100)  NOT NULL,
    prediction    VARCHAR(20)   NOT NULL,
    size          INTEGER  NOT NULL,
    odds          NUMERIC,
    result        VARCHAR(10)   NOT NULL,
    posted_at     TIMESTAMP     DEFAULT now()
);

CREATE TABLE esports_bets
(
    id            serial PRIMARY KEY,
    team          VARCHAR(100)  NOT NULL,
    prediction    VARCHAR(20)   NOT NULL,
    size          INTEGER  NOT NULL,
    odds          NUMERIC,
    result        VARCHAR(10)   NOT NULL,
    posted_at     TIMESTAMP     DEFAULT now()
);

CREATE TABLE open_bets
(
    id            serial PRIMARY KEY,
    message_id    VARCHAR(100)  NOT NULL,
    team          VARCHAR(100)  NOT NULL,
    prediction    VARCHAR(20)   NOT NULL,
    size          INTEGER  NOT NULL,
    odds          NUMERIC
);

CREATE TABLE users
(
    id            serial PRIMARY KEY,
    alias         VARCHAR(20) NOT NULL,
    images        jsonb
);

CREATE TABLE pets
(
    id            serial PRIMARY KEY,
    alias         VARCHAR(20) NOT NULL,
    images       jsonb
);

CREATE TABLE emotes
(
    id            serial PRIMARY KEY,
    alias         VARCHAR(20) NOT NULL,
    images        jsonb
);

CREATE INDEX idx_bets_date ON bets (posted_at);
CREATE INDEX idx_polo_bets_date ON polo_bets (posted_at);
CREATE INDEX idx_nick_bets_date ON nick_bets (posted_at);
CREATE INDEX idx_panos_bets_date ON panos_bets (posted_at);
CREATE INDEX idx_esports_bets_date ON esports_bets (posted_at);
CREATE INDEX idx_users ON users (alias);
CREATE INDEX idx_pets ON pets (alias);
CREATE INDEX idx_emotes ON emotes (alias);
