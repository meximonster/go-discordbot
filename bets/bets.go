package bets

import (
	"regexp"
	"strings"
	"time"

	"github.com/glenn-brown/golang-pkg-pcre/src/pkg/pcre"
)

var (
	betRegexp1       = regexp.MustCompile(`(.*?)((o|over|u|under)[0-9]{1,2}([.]5)?|X|x|1|2|1X|1x|2x|2X|X2|x2) [0-9]{1,3}u(.*)`)
	betRegexp2       = regexp.MustCompile(`(.*?)[0-9]{1,3}u ((o|over|u|under)[0-9]{1,2}([.]5)?|X|x|1|2|1X|1x|2x|2X|X2|x2)(.*)`)
	betRegexp3       = regexp.MustCompile(`(.*?)((o|over|u|under)[0-9]{1,2}([.]5)?|X|x|1|2|1X|1x|2x|2X|X2|x2)(.*?)[0-9]{1,3}u(.*)`)
	betRegexp4       = regexp.MustCompile(`(.*?)[0-9]{1,3}u(.*?)((o|over|u|under)[0-9]{1,2}([.]5)?|X|x|1|2|1X|1x|2x|2X|X2|x2)(.*)`)
	unitsRegexp      = regexp.MustCompile(`^[0-9]{1,3}u(.*?)$`)
	predictionRegexp = regexp.MustCompile(`^((o|over|u|under)[0-9]{1,2}([.]5)?|X|x|1|2|1X|1x|2x|2X|X2|x2)$`)
	oddsRegexp       = regexp.MustCompile(`^@([0-9]*[.])?[0-9]+$`)
	goalRegexp       = pcre.MustCompile(`([0-9])\1{2,}$`, 0)
)

type Bet struct {
	Id         string
	Team       string
	Prediction string
	Size       string
	Odds       string
	Result     string
	Created_at time.Time
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

func IsGoal(content string) bool {
	return goalRegexp.MatcherString(content, 0).MatchString(content, 0)
}

func Decouple(content string, result string) *Bet {
	bet := &Bet{}
	words := strings.Split(content, " ")
	var team string
	bet.Result = result
	for _, s := range words {

		if IsOdds(s) {
			bet.Odds = strings.Replace(s, "@", "", 1)
			continue
		}

		if IsPrediction(s) {
			bet.Prediction = s
			continue
		}

		if IsUnits(s) {
			bet.Size = s
			continue
		}

		// all regex checks failed, so the word must be part of the team name.
		if team == "" {
			team = team + s
		} else {
			team = team + " " + s
		}
	}

	bet.Team = team

	if bet.Team == "" || bet.Size == "" || bet.Prediction == "" {
		return &Bet{}
	}

	return bet
}
