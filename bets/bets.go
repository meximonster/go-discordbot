package bets

import (
	"regexp"
	"time"

	"github.com/glenn-brown/golang-pkg-pcre/src/pkg/pcre"
)

var (
	betRegexp1  = regexp.MustCompile(`(.*?)((o|over|u|under)[0-9]{1,2}([.]5)?|X|x) [0-9]{1,3}u(.*)`)
	betRegexp2  = regexp.MustCompile(`(.*?)[0-9]{1,3}u ((o|over|u|under)[0-9]{1,2}([.]5)?|X|x)(.*)`)
	betRegexp3  = regexp.MustCompile(`(.*?)((o|over|u|under)[0-9]{1,2}([.]5)?|X|x)(.*?)[0-9]{1,3}u(.*)`)
	betRegexp4  = regexp.MustCompile(`(.*?)[0-9]{1,3}u(.*?)((o|over|u|under)[0-9]{1,2}([.]5)?|X|x)(.*)`)
	unitsRegexp = regexp.MustCompile(`^[0-9]{1,3}u(.*?)`)
	goalRegexp  = pcre.MustCompile(`([0-9])\1{2,}$`, 0)
)

type Bet struct {
	Id         string
	Team       string
	Prediction string
	Size       string
	Result     string
	Created_at time.Time
}

func IsBet(content string) bool {
	return betRegexp1.MatchString(content) || betRegexp2.MatchString(content) || betRegexp3.MatchString(content) || betRegexp4.MatchString(content)
}

func IsUnits(word string) bool {
	return unitsRegexp.MatchString(word)
}

func IsGoal(content string) bool {
	return goalRegexp.MatcherString(content, 0).MatchString(content, 0)
}
