CREATE TABLE bets
(
    id            serial PRIMARY KEY,
    team          VARCHAR(100)  NOT NULL,
    prediction    VARCHAR(20)   NOT NULL,
    size          VARCHAR(100)  NOT NULL,
    result        VARCHAR(10)   NOT NULL,
    created_at    TIMESTAMP     DEFAULT NOW()
);

CREATE INDEX idx_bets_date ON bets (created_at);

