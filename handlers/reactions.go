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
		if isBet(m.Content) {
			switch r.Emoji.Name {
			case "✅":
				s.ChannelMessageSend(messageConfig.ChannelID, fmt.Sprintf("*** %s ----> WON!", m.Content))
			case "❌":
				s.ChannelMessageSend(messageConfig.ChannelID, fmt.Sprintf("*** %s *** ----> lost", m.Content))
			}
		}
	}
}
