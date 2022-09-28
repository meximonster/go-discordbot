package bet

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

var dbC *sqlx.DB

func NewDB(db *sqlx.DB) {
	dbC = db
}

func Ping() error {
	return dbC.Ping()
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

func GetUnitsPerMonth(table string) ([]UnitsPerMonth, error) {
	q := parseGraphQuery(unitPerMonthQuery, table)
	r := []UnitsPerMonth{}
	err := dbC.Select(&r, q)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func GetBetsPerMonth(table string) ([]BetsPerMonth, error) {
	q := parseGraphQuery(betsPerMonthQuery, table)
	r := []BetsPerMonth{}
	err := dbC.Select(&r, q)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func GetPercentBySize(table string) ([]PercentPerSize, error) {
	q := parseGraphQuery(percentPerSizeQuery, table)
	r := []PercentPerSize{}
	err := dbC.Select(&r, q)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func GetWonPerType(query string, table string) ([]float64, error) {
	q := parseGraphQuery(query, table)
	r := make([]float64, 0, 2)
	err := dbC.Select(&r, q)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func GetCountBySize(table string) ([]CountBySize, error) {
	q := parseGraphQuery(countBySizeQuery, table)
	r := []CountBySize{}
	err := dbC.Select(&r, q)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func GetCountByType(table string) ([]CountByType, error) {
	q := parseGraphQuery(countByTypeQuery, table)
	r := []CountByType{}
	err := dbC.Select(&r, q)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func GetYield(table string) ([]Yield, error) {
	q := parseGraphQuery(yieldQuery, table)
	var yield []Yield
	err := dbC.Select(&yield, q)
	if err != nil {
		return nil, err
	}
	return yield, nil
}
