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
	checkForBet(m.ChannelID, m.Author.ID, m.Content, s)
}

func checkForBet(channel string, author string, content string, s *discordgo.Session) {
	// Message was sent to pad-bets channel.
	if channel == messageConfig.ChannelID {
		// Author is pad.
		if author == messageConfig.UserID {
			if betRegexp1.MatchString(content) || betRegexp2.MatchString(content) || betRegexp3.MatchString(content) || betRegexp4.MatchString(content) {
				words := strings.Split(content, " ")
				for i := range words {
					if unitsRegexp.MatchString(words[i]) {
						betSizeStr := words[i][:strings.IndexByte(words[i], 'u')]
						betSizeInt, err := strconv.Atoi(betSizeStr)
						if err != nil {
							fmt.Println("error converting betSize to int: ", err)
							return
						}
						s.ChannelMessageSend(messageConfig.ChannelID, fmt.Sprintf("@everyone possible bet with %du stake was just posted.", betSizeInt))
					}
				}
			}
		}
	}
}

func checkAndReact(m *discordgo.MessageCreate, s *discordgo.Session) {
	content := strings.ToLower(m.Content)

	// Check for goal.
	if m.ChannelID == messageConfig.ChannelID && goalRegexp.MatcherString(m.Content, 0).MatchString(m.Content, 0) {
		s.ChannelMessageSend(messageConfig.ChannelID, "GOOOOOOOAAAAAAAAAAAAAAAALLLLL !!!!")
	}

	// Check for messages related to covid.
	if strings.Contains(content, "corona") || strings.Contains(content, "korona") || strings.Contains(content, "covid") {
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
	if strings.Contains(content, "kouvas") || strings.Contains(content, "κουβας") || strings.Contains(content, "κουβά") {
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
	if strings.Contains(content, "panagia") || strings.Contains(content, "παναγία") || strings.Contains(content, "παναγια") {
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

func ReactionCreate(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
	if r.ChannelID == messageConfig.ChannelID && (r.Emoji.Name == "✅" || r.Emoji.Name == "❌") {
		m, err := s.ChannelMessage(r.ChannelID, r.MessageID)
		if err != nil {
			fmt.Println("error getting message from reaction: ", err)
		}
		if betRegexp1.MatchString(m.Content) || betRegexp2.MatchString(m.Content) || betRegexp3.MatchString(m.Content) || betRegexp4.MatchString(m.Content) {
			if r.Emoji.Name == "✅" {
				s.ChannelMessageSend(messageConfig.ChannelID, fmt.Sprintf("*** %s ----> considered WON! ***", m.Content))
			} else {
				s.ChannelMessageSend(messageConfig.ChannelID, fmt.Sprintf("*** %s *** ----> considered lost ***", m.Content))
			}
		}
	}
}
