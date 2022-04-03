package handlers

import (
	"fmt"
	"regexp"

	"github.com/bwmarrin/discordgo"
	"github.com/glenn-brown/golang-pkg-pcre/src/pkg/pcre"
)

var (
	betRegexp1  = regexp.MustCompile(`(.*?)((o|over|u|under)[0-9]{1,2}([.]5)?|X|x) [0-9]{1,3}u(.*)`)
	betRegexp2  = regexp.MustCompile(`(.*?)[0-9]{1,3}u ((o|over|u|under)[0-9]{1,2}([.]5)?|X|x)(.*)`)
	betRegexp3  = regexp.MustCompile(`(.*?)((o|over|u|under)[0-9]{1,2}([.]5)?|X|x)(.*?)[0-9]{1,3}u(.*)`)
	betRegexp4  = regexp.MustCompile(`(.*?)[0-9]{1,3}u(.*?)((o|over|u|under)[0-9]{1,2}([.]5)?|X|x)(.*)`)
	goalRegexp  = pcre.MustCompile(`([0-9])\1{2,}$`, 0)
	unitsRegexp = regexp.MustCompile(`^[0-9]{1,3}u(.*?)`)
)

func isBet(content string) bool {
	return betRegexp1.MatchString(content) || betRegexp2.MatchString(content) || betRegexp3.MatchString(content) || betRegexp4.MatchString(content)
}

func respondWithImage(channel string, title string, imageURL string, s *discordgo.Session) {
	_, err := s.ChannelMessageSendEmbed(channel, &discordgo.MessageEmbed{
		Title: title,
		Image: &discordgo.MessageEmbedImage{
			URL: imageURL,
		},
	})
	if err != nil {
		fmt.Println("error sending image: ", err)
	}
}
