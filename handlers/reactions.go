package handlers

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/meximonster/go-discordbot/bets"
)

func ReactionCreate(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
	if r.ChannelID == messageConfig.ChannelID && (r.Emoji.Name == "✅" || r.Emoji.Name == "❌") {
		m, err := s.ChannelMessage(r.ChannelID, r.MessageID)
		if err != nil {
			fmt.Println("error getting message from reaction: ", err)
		}
		if bets.IsBet(m.Content) {
			var b *bets.Bet
			switch r.Emoji.Name {
			case "✅":
				s.ChannelMessageSend(messageConfig.ChannelID, fmt.Sprintf("*** %s ----> WON!", m.Content))
				b = bets.Decouple(m.Content, "won")
			case "❌":
				s.ChannelMessageSend(messageConfig.ChannelID, fmt.Sprintf("*** %s ----> lost", m.Content))
				b = bets.Decouple(m.Content, "lost")
			}
			s.ChannelMessageSend(messageConfig.ChannelID, fmt.Sprintf("BET INFO: Team: %s, Prediction: %s, Size: %s, Odds: %s, Result: %s", b.Team, b.Prediction, b.Size, b.Odds, b.Result))
		}
	}
}
