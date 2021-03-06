package handlers

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	bet "github.com/meximonster/go-discordbot/bet"
)

func ReactionCreate(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
	if (r.ChannelID == betMsgConf.ChannelID || r.ChannelID == poloMsgConf.ChannelID) && (r.Emoji.Name == "✅" || r.Emoji.Name == "❌") {
		m, err := s.ChannelMessage(r.ChannelID, r.MessageID)
		if err != nil {
			s.ChannelMessageSend(r.ChannelID, err.Error())
			return
		}
		// Ignore reactions to bot messages.
		if m.Author.ID == s.State.User.ID {
			return
		}
		if bet.IsBet(m.Content) {
			var result string
			switch r.Emoji.Name {
			case "✅":
				s.ChannelMessageSend(r.ChannelID, fmt.Sprintf("***"+"%s ----> WON!"+"***", m.Content))
				result = "won"
			case "❌":
				s.ChannelMessageSend(r.ChannelID, fmt.Sprintf("***"+"%s ----> lost"+"***", m.Content))
				result = "lost"
			}
			table := tableRef(r.ChannelID)
			b, err := bet.Decouple(m.Content, result, table)
			if err != nil {
				s.ChannelMessageSend(r.ChannelID, err.Error())
				return
			}
			err = bet.Store(b, table)
			if err != nil {
				s.ChannelMessageSend(r.ChannelID, err.Error())
				return
			}
		}
	}
}
