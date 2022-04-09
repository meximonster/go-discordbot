package bets

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

var dbC *sqlx.DB

func NewDB(db *sqlx.DB) {
	dbC = db
}

func CloseDB() {
	err := dbC.Close()
	if err != nil {
		fmt.Println(err)
	}
}

func (b *Bet) Store(table string) error {
	q := `INSERT INTO $1 (team,prediction,size,odds,result) VALUES ($2,$3,$4,$5,$6)`
	_ = dbC.MustExec(q, table, b.Team, b.Prediction, b.Size, b.Odds, b.Result)
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
