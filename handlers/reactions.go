package handlers

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

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
