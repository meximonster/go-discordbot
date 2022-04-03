package handlers

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/meximonster/go-discordbot/bets"
)

var messageConfig *MessageInfo

type MessageInfo struct {
	ChannelID string
	UserID    string
}

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

	checkAndRespond(m, s)
	checkForBet(m.ChannelID, m.Author.ID, m.Content, s)
}

func checkForBet(channel string, author string, content string, s *discordgo.Session) {
	// Message was sent to pad-bets channel.
	if channel == messageConfig.ChannelID {
		// Author is pad.
		if author == messageConfig.UserID {
			if bets.IsBet(content) {
				words := strings.Split(content, " ")
				for i := range words {
					if bets.IsUnits(words[i]) {
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

func checkAndRespond(m *discordgo.MessageCreate, s *discordgo.Session) {
	content := strings.ToLower(m.Content)

	// return repo url.
	if m.Content == "!git" {
		s.ChannelMessageSend(m.ChannelID, "https://github.com/meximonster/go-discordbot")
	}

	// Check for goal.
	if m.ChannelID == messageConfig.ChannelID && bets.IsGoal(m.Content) {
		s.ChannelMessageSend(messageConfig.ChannelID, "GOOOOOOOAAAAAAAAAAAAAAAALLLLL !!!!")
	}

	// Check for messages related to aalesund.
	if strings.Contains(content, "alesund") {
		s.ChannelMessageSend(m.ChannelID, ":sweat_drops:")
	}

	// Check for messages related to begging for something.
	if strings.Contains(content, "please") || strings.Contains(content, "plz") || strings.Contains(content, "pliz") {
		s.MessageReactionAdd(m.ChannelID, m.ID, "🙏")
	}

	// Check for messages related to covid.
	if strings.Contains(content, "corona") || strings.Contains(content, "korona") || strings.Contains(content, "covid") {
		respondWithImage(m.ChannelID, "covid ????", "https://pbs.twimg.com/ext_tw_video_thumb/1239694832781512705/pu/img/zKpSNMMa_-8d9bFo.jpg", s)
	}

	// Check for messages related to kouvas.
	if strings.Contains(content, "kouvas") || strings.Contains(content, "κουβας") || strings.Contains(content, "κουβά") {
		respondWithImage(m.ChannelID, "mia zwh kouvas", "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTYLsQwNnLkEyL1MeAAegoEJDs8KOYE6AtXng&usqp=CAU", s)
	}

	// Check for messages related to panagia.
	if strings.Contains(content, "panagia") || strings.Contains(content, "παναγία") || strings.Contains(content, "παναγια") {
		respondWithImage(m.ChannelID, "gamw thn panagia", "https://www.in.gr/wp-content/uploads/2019/08/23.png", s)
	}
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
