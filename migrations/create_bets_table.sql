CREATE TABLE bets
(
    id            serial PRIMARY KEY,
    team          VARCHAR(100)  NOT NULL,
    prediction    VARCHAR(20)   NOT NULL,
    size          INTEGER  NOT NULL,
    odds          NUMERIC,
    result        VARCHAR(10)   NOT NULL,
    posted_at     TIMESTAMP     DEFAULT NOW()
);

CREATE INDEX idx_bets_date ON bets (posted_at);

