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

CREATE INDEX idx_polo_bets_date ON polo_bets (posted_at);

