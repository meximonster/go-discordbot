package handlers

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	bet "github.com/meximonster/go-discordbot/bet"
)

func ReactionCreate(s *discordgo.Session, r *discordgo.MessageReactionAdd) {

	if !(r.Emoji.Name == "✅" || r.Emoji.Name == "❌") {
		return
	}

	if !bet.IsBetChannel(r.ChannelID) {
		return
	}

	m, err := s.ChannelMessage(r.ChannelID, r.MessageID)
	if err != nil {
		s.ChannelMessageSend(r.ChannelID, err.Error())
		return
	}

	// Ignore reactions to bot messages.
	if m.Author.ID == s.State.User.ID {
		return
	}

	if !bet.IsAdmin(m.Author.ID) {
		return
	}

	if !bet.IsBet(m.Content) {
		return
	}

	var result string
	switch r.Emoji.Name {
	case "✅":
		s.ChannelMessageSend(r.ChannelID, fmt.Sprintf("***"+"%s ✅"+"***", m.Content))
		result = "won"
	case "❌":
		s.ChannelMessageSend(r.ChannelID, fmt.Sprintf("***"+"%s ❌"+"***", m.Content))
		result = "lost"
	}

	table := bet.GetTableFromChannel(r.ChannelID)
	b, err := bet.Decouple(m.Content, result)
	if err != nil {
		s.ChannelMessageSend(r.ChannelID, err.Error())
		return
	}

	if b.Odds == 0 {
		s.ChannelMessageSend(r.ChannelID, "settling bet with null odds. fix the odds and try again")
		return
	}

	err = bet.Store(b, table)
	if err != nil {
		s.ChannelMessageSend(r.ChannelID, err.Error())
		return
	}

	bet.Settle(m.ID)

}
