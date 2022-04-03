package handlers

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/glenn-brown/golang-pkg-pcre/src/pkg/pcre"
)

var messageConfig *MessageInfo

type MessageInfo struct {
	ChannelID string
	UserID    string
}

var (
	betRegexp1  = regexp.MustCompile(`(.*?)(o|over|u|under)[0-9]{1,2}([.]5)? [0-9]{1,3}u(.*)`)
	betRegexp2  = regexp.MustCompile(`(.*?)[0-9]{1,3}u (o|over|u|under)[0-9]{1,2}([.]5)?(.*)`)
	betRegexp3  = regexp.MustCompile(`(.*?)(o|over|u|under)[0-9]{1,2}([.]5)?(.*?)[0-9]{1,3}u(.*)`)
	betRegexp4  = regexp.MustCompile(`(.*?)[0-9]{1,3}u(.*?)(o|over|u|under)[0-9]{1,2}([.]5)?(.*)`)
	goalRegexp  = pcre.MustCompile(`([0-9])\1{2,}$`, 0)
	unitsRegexp = regexp.MustCompile(`^[0-9]{1,3}u(.*?)`)
)

func MessageConfigInit(channel string, user string) {
	messageConfig = &MessageInfo{
		ChannelID: channel,
		UserID:    user,
	}
}

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself.
	if m.Author.ID == s.State.User.ID {
		return
	}

	checkAndReact(m, s)
	checkForBet(m, s)
}

func checkForBet(m *discordgo.MessageCreate, s *discordgo.Session) {
	// Message was sent to pad-bets channel.
	if m.ChannelID == messageConfig.ChannelID {
		// Author is pad.
		if m.Author.ID == messageConfig.UserID {
			if betRegexp1.MatchString(m.Content) || betRegexp2.MatchString(m.Content) || betRegexp3.MatchString(m.Content) || betRegexp4.MatchString(m.Content) {
				words := strings.Split(m.Content, " ")
				for i := range words {
					if unitsRegexp.MatchString(words[i]) {
						betSizeStr := words[i][:strings.IndexByte(words[i], 'u')]
						betSizeInt, err := strconv.Atoi(betSizeStr)
						if err != nil {
							fmt.Println("error converting betSize to int: ", err)
							return
						}
						if betSizeInt >= 15 {
							s.ChannelMessageSend(messageConfig.ChannelID, fmt.Sprintf("@everyone possible bet with %du stake was just posted.", betSizeInt))
						}
					}
				}
			}
		}
	}
}

func checkAndReact(m *discordgo.MessageCreate, s *discordgo.Session) {
	// Check for goal.
	if m.ChannelID == messageConfig.ChannelID && goalRegexp.MatcherString(m.Content, 0).MatchString(m.Content, 0) {
		s.ChannelMessageSend(messageConfig.ChannelID, "GOOOOOOOAAAAAAAAAAAAAAAALLLLL !!!!")
	}

	// Check for messages related to covid.
	if strings.Contains(m.Content, "corona") || strings.Contains(m.Content, "korona") || strings.Contains(m.Content, "covid") {
		_, err := s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
			Title: "covid???",
			Image: &discordgo.MessageEmbedImage{
				URL: "https://pbs.twimg.com/ext_tw_video_thumb/1239694832781512705/pu/img/zKpSNMMa_-8d9bFo.jpg",
			},
		})
		if err != nil {
			fmt.Println("error sending image: ", err)
		}
	}

	// Check for messages related to kouvas.
	if strings.Contains(m.Content, "kouvas") || strings.Contains(m.Content, "κουβας") || strings.Contains(m.Content, "κουβά") {
		_, err := s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
			Title: "kouvas",
			Image: &discordgo.MessageEmbedImage{
				URL: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTYLsQwNnLkEyL1MeAAegoEJDs8KOYE6AtXng&usqp=CAU",
			},
		})
		if err != nil {
			fmt.Println("error sending image: ", err)
		}
	}

	// Check for messages related to panagia.
	if strings.Contains(m.Content, "panagia") || strings.Contains(m.Content, "παναγία") || strings.Contains(m.Content, "παναγια") {
		_, err := s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
			Title: "gamw thn panagia",
			Image: &discordgo.MessageEmbedImage{
				URL: "https://www.in.gr/wp-content/uploads/2019/08/23.png",
			},
		})
		if err != nil {
			fmt.Println("error sending image: ", err)
		}
	}
}
