package handlers

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/meximonster/go-discordbot/bets"
)

func ReactionCreate(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
	if (r.ChannelID == padMsgConf.ChannelID || r.ChannelID == fykMsgConf.ChannelID) && (r.Emoji.Name == "✅" || r.Emoji.Name == "❌") {
		m, err := s.ChannelMessage(r.ChannelID, r.MessageID)
		if err != nil {
			fmt.Println("error getting message from reaction: ", err)
			return
		}
		if bets.IsBet(m.Content) {
			var result string
			var table string
			switch r.Emoji.Name {
			case "✅":
				s.ChannelMessageSend(r.ChannelID, fmt.Sprintf("***"+"%s ----> WON!"+"***", m.Content))
				result = "won"
			case "❌":
				s.ChannelMessageSend(r.ChannelID, fmt.Sprintf("***"+"%s ----> lost"+"***", m.Content))
				result = "lost"
			}
			if r.ChannelID == padMsgConf.ChannelID {
				table = "bets"
			} else {
				table = "polo_bets"
			}
			b, err := bets.DecoupleAndStore(m.Content, result, table)
			if err != nil {
				s.ChannelMessageSend(r.ChannelID, err.Error())
				return
			}
			s.ChannelMessageSend(r.ChannelID, fmt.Sprintf("BET INFO: Team: *%s*, Prediction: *%s*, Size: *%d*, Odds: *%v*, Result: *%s*", b.Team, b.Prediction, b.Size, b.Odds, b.Result))
		}
	}
}
