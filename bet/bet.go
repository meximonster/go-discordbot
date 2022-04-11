package bet

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	betRegexp1       = regexp.MustCompile(`(.*?)((o|over|u|under|\+|\-)[0-9]{1,2}([.]7?5)?|X|x|1|2|1X|1x|2x|2X|X2|x2) [0-9]{1,3}u(.*)`)
	betRegexp2       = regexp.MustCompile(`(.*?)[0-9]{1,3}u ((o|over|u|under|\+|\-)[0-9]{1,2}([.]7?5)?|X|x|1|2|1X|1x|2x|2X|X2|x2)(.*)`)
	betRegexp3       = regexp.MustCompile(`(.*?)((o|over|u|under|\+|\-)[0-9]{1,2}([.]7?5)?|X|x|1|2|1X|1x|2x|2X|X2|x2)(.*?)[0-9]{1,3}u(.*)`)
	betRegexp4       = regexp.MustCompile(`(.*?)[0-9]{1,3}u(.*?)((o|over|u|under|\+|\-)[0-9]{1,2}([.]7?5)?|X|x|1|2|1X|1x|2x|2X|X2|x2)(.*)`)
	unitsRegexp      = regexp.MustCompile(`^[0-9]{1,3}u(.*?)$`)
	predictionRegexp = regexp.MustCompile(`^((o|over|u|under|\+|\-)[0-9]{1,2}([.]7?5)?|X|x|1|2|1X|1x|2x|2X|X2|x2|combo|[0-9]ada)$`)
	oddsRegexp       = regexp.MustCompile(`^@([0-9]*[.])?[0-9]+$`)
)

type Bet struct {
	Id         string
	Team       string
	Prediction string
	Size       int
	Odds       float64
	Result     string
	Posted_at  time.Time
}

type BetSummary struct {
	Count       int
	Total_units int
	Result      string
}

func IsBet(content string) bool {
	return betRegexp1.MatchString(content) || betRegexp2.MatchString(content) || betRegexp3.MatchString(content) || betRegexp4.MatchString(content)
}

func IsPrediction(word string) bool {
	return predictionRegexp.MatchString(word)
}

func IsOdds(word string) bool {
	return oddsRegexp.MatchString(word)
}

func IsUnits(word string) bool {
	return unitsRegexp.MatchString(word)
}

func Decouple(content string, result string, table string) (Bet, error) {
	var b Bet
	words := strings.Split(content, " ")
	var team string
	for _, s := range words {

		if IsOdds(s) {
			o := strings.Replace(s, "@", "", 1)
			fl, err := strconv.ParseFloat(o, 64)
			if err != nil {
				return Bet{}, fmt.Errorf("error parsing float: %s", err.Error())
			}
			b.Odds = fl
			continue
		}

		if IsPrediction(s) {
			b.Prediction = s
			continue
		}

		if IsUnits(s) {
			o := strings.Replace(s, "u", "", 1)
			i, err := strconv.Atoi(o)
			if err != nil {
				return Bet{}, fmt.Errorf("error parsing int: %s", err.Error())
			}
			b.Size = i
			continue
		}

		// all regex checks failed, so the word must be part of the team name.
		if team == "" {
			team = team + s
		} else {
			team = team + " " + s
		}
	}

	b.Team = team
	b.Result = result

	if b.Team == "" || b.Prediction == "" || b.Size == 0 {
		return b, fmt.Errorf("discarding bet: INFO: Team: %s, Prediction: %s, Size: %d", b.Team, b.Prediction, b.Size)
	}

	return b, nil
}

func Store(b Bet, table string) error {
	err := b.Store(table)
	if err != nil {
		return err
	}
	return nil
}

func (b *Bet) Format() string {
	return fmt.Sprintf("%s %s %du ---> %s\n", b.Team, b.Prediction, b.Size, b.Result)
}

func (bs *BetSummary) Format() string {
	return fmt.Sprintf("%d bets %s, total_units: %d\n", bs.Count, bs.Result, bs.Total_units)
}

func FormatBets(bets []Bet) string {
	betFormats := make([]string, len(bets))
	for i, b := range bets {
		betFormats[i] = b.Format()
	}
	var result string
	for i := range betFormats {
		result = result + betFormats[i]
	}
	return result
}

func FormatBetsSum(sum []BetSummary) string {
	sumFormats := make([]string, len(sum))
	var net int
	for i, s := range sum {
		if s.Result == "won" {
			net += s.Total_units
		} else {
			net -= s.Total_units
		}
		sumFormats[i] = s.Format()
	}
	var result string
	for i := range sumFormats {
		result = result + sumFormats[i]
	}
	result = result + fmt.Sprintf("profit/loss: %d", net)
	return result
}
