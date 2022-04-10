package handlers

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/bwmarrin/discordgo"
	bet "github.com/meximonster/go-discordbot/bet"
	"github.com/meximonster/go-discordbot/queries"
)

var (
	padMsgConf      *MessageInfo
	fykMsgConf      *MessageInfo
	parolaChannelID string
)

type MessageInfo struct {
	ChannelID string
	UserID    string
}

func MessageConfigInit(padChannel string, padID string, fykChannel string, fykID string, parolaChannel string) {
	padMsgConf = &MessageInfo{
		ChannelID: padChannel,
		UserID:    padID,
	}
	fykMsgConf = &MessageInfo{
		ChannelID: fykChannel,
		UserID:    fykID,
	}
	parolaChannelID = parolaChannel
}

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself.
	if m.Author.ID == s.State.User.ID {
		return
	}

	checkAndRespond(m, s)
	checkForParola(m, s)
	checkForBet(m.ChannelID, m.Author.ID, m.Content, s)
	checkForBetQuery(m, s)
	checkForBetSumQuery(m, s)
}

func checkForParola(m *discordgo.MessageCreate, s *discordgo.Session) {
	if m.ChannelID == parolaChannelID {
		if len(m.Attachments) > 0 || strings.HasPrefix(m.Content, "https://www.stoiximan.gr/mybets/") {
			s.ChannelMessageSend(m.ChannelID, "@everyone possible parola was just posted.")
		}
	}
}

func checkForBet(channel string, author string, content string, s *discordgo.Session) {
	if (channel == padMsgConf.ChannelID && author == padMsgConf.UserID) || (channel == fykMsgConf.ChannelID && author == fykMsgConf.UserID) {
		if bet.IsBet(content) {
			table := tableRef(channel)
			b, err := bet.Decouple(content, "", table)
			if err != nil {
				s.ChannelMessageSend(channel, err.Error())
				return
			}
			s.ChannelMessageSend(channel, fmt.Sprintf("%s %s %du @everyone", b.Team, b.Prediction, b.Size))
		}
	}
}

func checkAndRespond(m *discordgo.MessageCreate, s *discordgo.Session) {
	content := strings.ToLower(m.Content)

	// return repo url.
	if m.Content == "!git" {
		s.ChannelMessageSend(m.ChannelID, "https://github.com/meximonster/go-discordbot")
	}

	if m.Content == "!giannakis" {
		rng := rand.Intn(10)
		var image, text string
		if rng < 3 {
			image = "https://i.imgur.com/VocVxhr.jpg"
			text = "mpainei ez"
		} else if rng >= 3 && rng < 7 {
			image = "https://i.imgur.com/yBw8qEU.jpg"
			text = "eixame"
		} else {
			image = "https://i.imgur.com/vfyPcEB.jpg"
			text = "irtha kai to vazw"
		}
		respondWithImage(m.ChannelID, text, image, s)
	}

	// Check for goal.
	if m.ChannelID == padMsgConf.ChannelID && bet.IsGoal(m.Content) {
		s.ChannelMessageSend(padMsgConf.ChannelID, "GOOOOOOOAAAAAAAAAAAAAAAALLLLL !!!!")
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
		respondWithImage(m.ChannelID, "covid ????", "https://i.imgur.com/Ydm7d7l.jpg", s)
	}

	// Check for messages related to kouvas.
	if strings.Contains(content, "kouvas") || strings.Contains(content, "κουβας") || strings.Contains(content, "κουβά") {
		respondWithImage(m.ChannelID, "mia zwh kouvas", "https://i.imgur.com/XccIGz2.jpg", s)
	}

	// Check for messages related to panagia.
	if strings.Contains(content, "panagia") || strings.Contains(content, "παναγία") || strings.Contains(content, "παναγια") {
		respondWithImage(m.ChannelID, "gamw thn panagia", "https://i.imgur.com/eypNquJ.png", s)
	}
}

func checkForBetQuery(m *discordgo.MessageCreate, s *discordgo.Session) {
	if (m.ChannelID == padMsgConf.ChannelID || m.ChannelID == fykMsgConf.ChannelID) && strings.HasPrefix(m.Content, "!bet ") {
		table := tableRef(m.ChannelID)
		q := queries.Parse(m.Content, table)
		bets, err := bet.GetBetsByQuery(q)
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("error getting bets: %s", err.Error()))
			return
		}
		if len(bets) == 0 {
			s.ChannelMessageSend(m.ChannelID, "no results")
			return
		}
		res := bet.FormatBets(bets)
		s.ChannelMessageSend(m.ChannelID, res)
	}
}

func checkForBetSumQuery(m *discordgo.MessageCreate, s *discordgo.Session) {
	if (m.ChannelID == padMsgConf.ChannelID || m.ChannelID == fykMsgConf.ChannelID) && strings.HasPrefix(m.Content, "!betsum ") {
		table := tableRef(m.ChannelID)
		q := queries.ParseSum(m.Content, table)
		sum, err := bet.GetBetsSumByQuery(q)
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("error getting bets: %s", err.Error()))
			return
		}
		if len(sum) == 0 {
			s.ChannelMessageSend(m.ChannelID, "no results")
			return
		}
		res := bet.FormatBetsSum(sum)
		s.ChannelMessageSend(m.ChannelID, res)
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

func tableRef(channel string) string {
	var table string
	if channel == padMsgConf.ChannelID {
		table = "bets"
	} else {
		table = "polo_bets"
	}
	return table
}
