package bets

import (
	"github.com/jmoiron/sqlx"
)

var dbC *sqlx.DB

func NewDB(db *sqlx.DB) {
	dbC = db
}

func (b *Bet) Store() error {
	q := `INSERT INTO bets (team,prediction,size,odds,result) VALUES ($1,$2,$3,$4,$5)`
	_ = dbC.MustExec(q, b.Team, b.Prediction, b.Size, b.Odds, b.Result)
	return nil
}

func GetByQuery(query string) ([]Bet, error) {
	bets := []Bet{}
	err := dbC.Select(&bets, query)
	if err != nil {
		return nil, err
	}
	return bets, nil
}
