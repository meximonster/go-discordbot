package handlers

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	bet "github.com/meximonster/go-discordbot/bet"
)

func ReactionCreate(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
	if (r.ChannelID == padMsgConf.ChannelID || r.ChannelID == fykMsgConf.ChannelID) && (r.Emoji.Name == "✅" || r.Emoji.Name == "❌") {
		m, err := s.ChannelMessage(r.ChannelID, r.MessageID)
		if err != nil {
			fmt.Println("error getting message from reaction: ", err)
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
			_ = bet.Store(b, table)
			s.ChannelMessageSend(r.ChannelID, fmt.Sprintf("BET INFO: Team: *%s*, Prediction: *%s*, Size: *%d*, Odds: *%v*, Result: *%s*", b.Team, b.Prediction, b.Size, b.Odds, b.Result))
		}
	}
}
