package bets

import (
	"fmt"

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

func (b *Bet) GetById(id int) (*Bet, error) {
	q := `SELECT id,team,prediction,size,odds,result,posted_at FROM bets`
	if err := dbC.Get(&b, q, id); err != nil {
		return &Bet{}, fmt.Errorf("error getting bet: %s", err.Error())
	}
	return b, nil
}
