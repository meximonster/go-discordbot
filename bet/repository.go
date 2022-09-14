package bet

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

var dbC *sqlx.DB

func NewDB(db *sqlx.DB) {
	dbC = db
}

func CloseDB() error {
	err := dbC.Close()
	if err != nil {
		return err
	}
	return nil
}

func (b *Bet) Store(table string) error {
	q := fmt.Sprintf(`INSERT INTO %s (team,prediction,size,odds,result) VALUES ($1,$2,$3,$4,$5)`, table)
	_, err := dbC.Exec(q, b.Team, b.Prediction, b.Size, b.Odds, b.Result)
	return err
}

func GetBetsByQuery(query string) ([]Bet, error) {
	bets := []Bet{}
	err := dbC.Select(&bets, query)
	if err != nil {
		return nil, err
	}
	return bets, nil
}

func GetBetsSumByQuery(query string) ([]BetSummary, error) {
	sum := []BetSummary{}
	err := dbC.Select(&sum, query)
	if err != nil {
		return nil, err
	}
	return sum, nil
}

func getUnitsPerMonth() ([]UnitsPerMonth, error) {
	r := []UnitsPerMonth{}
	err := dbC.Select(&r, unitPerMonthQuery)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func getBetsPerMonth() ([]BetsPerMonth, error) {
	r := []BetsPerMonth{}
	err := dbC.Select(&r, betsPerMonthQuery)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func getPercentBySize() ([]PercentPerSize, error) {
	r := []PercentPerSize{}
	err := dbC.Select(&r, percentPerSizeQuery)
	if err != nil {
		return nil, err
	}
	return r, nil
}
