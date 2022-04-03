package handlers

import (
	"regexp"

	"github.com/glenn-brown/golang-pkg-pcre/src/pkg/pcre"
)

var (
	betRegexp1  = regexp.MustCompile(`(.*?)(o|over|u|under)[0-9]{1,2}([.]5)? [0-9]{1,3}u(.*)`)
	betRegexp2  = regexp.MustCompile(`(.*?)[0-9]{1,3}u (o|over|u|under)[0-9]{1,2}([.]5)?(.*)`)
	betRegexp3  = regexp.MustCompile(`(.*?)(o|over|u|under)[0-9]{1,2}([.]5)?(.*?)[0-9]{1,3}u(.*)`)
	betRegexp4  = regexp.MustCompile(`(.*?)[0-9]{1,3}u(.*?)(o|over|u|under)[0-9]{1,2}([.]5)?(.*)`)
	goalRegexp  = pcre.MustCompile(`([0-9])\1{2,}$`, 0)
	unitsRegexp = regexp.MustCompile(`^[0-9]{1,3}u(.*?)`)
)

func isBet(content string) bool {
	return betRegexp1.MatchString(content) || betRegexp2.MatchString(content) || betRegexp3.MatchString(content) || betRegexp4.MatchString(content)
}
