package handlers

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	bet "github.com/meximonster/go-discordbot/bet"
	"github.com/meximonster/go-discordbot/configuration"
	"github.com/meximonster/go-discordbot/queries"
	"github.com/meximonster/go-discordbot/user"
)

var (
	padMsgConf      *MessageInfo
	fykMsgConf      *MessageInfo
	userNames       []string
	parolaChannelID string
	banlist         []string
)

type MessageInfo struct {
	ChannelID string
	UserID    string
}

func MessageConfigInit(users []configuration.UserConfig, parolaChannel string, blacklist []string) {
	parolaChannelID = parolaChannel
	for _, u := range users {
		switch strings.ToLower(u.Username) {
		case "pad":
			padMsgConf = &MessageInfo{
				UserID:    u.UserID,
				ChannelID: u.ChannelID,
			}
		case "fyk":
			fykMsgConf = &MessageInfo{
				UserID:    u.UserID,
				ChannelID: u.ChannelID,
			}
		}
		userNames = append(userNames, u.Username)
	}
	banlist = blacklist
}

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself.
	if m.Author.ID == s.State.User.ID {
		return
	}

	serveBanlist(m, s)
	checkAndRespond(m, s)
	checkForUser(m, s)
	checkForParola(m, s)
	checkForBet(m.ChannelID, m.Author.ID, m.Content, s)
	checkForBetQuery(m, s)
	checkForBetSumQuery(m, s)
}

func serveBanlist(m *discordgo.MessageCreate, s *discordgo.Session) {
	if m.Content == "!banlist" {
		var result string
		for i, banword := range banlist {
			result = result + fmt.Sprintf("%d. %s\n", i, banword)
		}
		s.ChannelMessageSend(m.ChannelID, result)
	}
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

func checkForUser(m *discordgo.MessageCreate, s *discordgo.Session) {
	str := strings.TrimPrefix(m.Content, "!")
	for _, uname := range userNames {
		if str == strings.ToLower(uname) {
			respondWithRandomImage(uname, m.ChannelID, s)
		}
	}
}

func checkAndRespond(m *discordgo.MessageCreate, s *discordgo.Session) {
	content := strings.ToLower(m.Content)

	// return repo url.
	if m.Content == "!git" {
		s.ChannelMessageSend(m.ChannelID, "https://github.com/meximonster/go-discordbot")
	}

	// Check for messages related to covid.
	if strings.Contains(content, "corona") || strings.Contains(content, "korona") || strings.Contains(content, "covid") {
		respondWithImage(m.ChannelID, "covid ????", "https://i.imgur.com/Ydm7d7l.jpg", s)
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

func respondWithRandomImage(name string, channel string, s *discordgo.Session) {
	u := user.GetUserByName(name)
	img := u.RandomImage()
	respondWithImage(channel, img.Text, img.Url, s)
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
