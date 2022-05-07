package handlers

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	bet "github.com/meximonster/go-discordbot/bet"
	"github.com/meximonster/go-discordbot/configuration"
	"github.com/meximonster/go-discordbot/meme"
	"github.com/meximonster/go-discordbot/queries"
	"github.com/meximonster/go-discordbot/user"
)

var (
	padMsgConf      *betMsgSrc
	fykMsgConf      *betMsgSrc
	userNames       []string
	parolaChannelID string
	banlist         []string
)

type betMsgSrc struct {
	ChannelID string
	UserID    string
}

func MessageConfigInit(users []configuration.UserConfig, parolaChannel string, blacklist []string) {
	parolaChannelID = parolaChannel
	for _, u := range users {
		switch strings.ToLower(u.Username) {
		case "pad":
			padMsgConf = &betMsgSrc{
				UserID:    u.UserID,
				ChannelID: u.ChannelID,
			}
		case "fyk":
			fykMsgConf = &betMsgSrc{
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

	if m.Content == "!git" {
		s.ChannelMessageSend(m.ChannelID, "https://github.com/meximonster/go-discordbot")
	}

	serveMeme(m.Content, m.ChannelID, s)
	serveBanlist(m.Content, m.ChannelID, s)
	serveUsers(m.Content, m.ChannelID, s)
	servePets(m.Content, m.ChannelID, s)
	checkForUser(m.Content, m.ChannelID, s)
	checkForParola(m.Content, m.ChannelID, m.Attachments, s)
	checkForBet(m.ChannelID, m.Author.ID, m.Content, s)
	checkForBetQuery(m.Content, m.ChannelID, s)
	checkForBetSumQuery(m.Content, m.ChannelID, s)
}

func serveUsers(content string, channel string, s *discordgo.Session) {
	if content == "!users" {
		users := user.GetUsers()
		if len(users) == 0 {
			s.ChannelMessageSend(channel, "no users configured")
			return
		}
		var str string
		cnt := 0
		for _, u := range users {
			str = str + fmt.Sprintf("%d. %s\n", cnt+1, u.Username)
			cnt++
		}
		result := "Configured users are:\n" + str
		s.ChannelMessageSend(channel, result)
	}
}

func servePets(content string, channel string, s *discordgo.Session) {
	if content == "!pets" {
		pets := user.GetPets()
		if len(pets) == 0 {
			s.ChannelMessageSend(channel, "no pets configured")
			return
		}
		var str string
		cnt := 0
		for _, u := range pets {
			str = str + fmt.Sprintf("%d. %s\n", cnt+1, u.Username)
			cnt++
		}
		result := "Configured pets are:\n" + str
		s.ChannelMessageSend(channel, result)
	}
}

func serveBanlist(content string, channel string, s *discordgo.Session) {
	if content == "!banlist" {
		var result string
		for i, banword := range banlist {
			result = result + fmt.Sprintf("%d. %s\n", i+1, banword)
		}
		s.ChannelMessageSend(channel, result)
	}
}

func serveMeme(content string, channel string, s *discordgo.Session) {
	if content == "!meme" {
		link, url, err := meme.Random()
		if err != nil {
			s.ChannelMessageSend(channel, err.Error())
			return
		}
		respondWithImage(channel, link, url, s)
	}
}

func checkForParola(content string, channel string, attachments []*discordgo.MessageAttachment, s *discordgo.Session) {
	if channel == parolaChannelID {
		if len(attachments) > 0 || strings.HasPrefix(content, "https://www.stoiximan.gr/mybets/") {
			s.ChannelMessageSend(channel, "@everyone possible parola was just posted.")
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

func checkForUser(content string, channel string, s *discordgo.Session) {
	if strings.HasPrefix(content, "!") {
		str := strings.TrimPrefix(content, "!")
		for _, uname := range userNames {
			if str == strings.ToLower(uname) {
				respondWithRandomImage(uname, channel, s)
			}
		}
	}
}

func checkForBetQuery(content string, channel string, s *discordgo.Session) {
	if (channel == padMsgConf.ChannelID || channel == fykMsgConf.ChannelID) && strings.HasPrefix(content, "!bet ") {
		table := tableRef(channel)
		q := queries.Parse(content, table)
		bets, err := bet.GetBetsByQuery(q)
		if err != nil {
			s.ChannelMessageSend(channel, fmt.Sprintf("error getting bets: %s", err.Error()))
			return
		}
		if len(bets) == 0 {
			s.ChannelMessageSend(channel, "no results")
			return
		}
		res := bet.FormatBets(bets)
		s.ChannelMessageSend(channel, res)
	}
}

func checkForBetSumQuery(content string, channel string, s *discordgo.Session) {
	if (channel == padMsgConf.ChannelID || channel == fykMsgConf.ChannelID) && strings.HasPrefix(content, "!betsum ") {
		table := tableRef(channel)
		q := queries.ParseSum(content, table)
		sum, err := bet.GetBetsSumByQuery(q)
		if err != nil {
			s.ChannelMessageSend(channel, fmt.Sprintf("error getting bets: %s", err.Error()))
			return
		}
		if len(sum) == 0 {
			s.ChannelMessageSend(channel, "no results")
			return
		}
		res := bet.FormatBetsSum(sum)
		s.ChannelMessageSend(channel, res)
	}
}

func respondWithRandomImage(name string, channel string, s *discordgo.Session) {
	u, err := user.GetByName(name)
	if err != nil {
		s.ChannelMessageSend(channel, err.Error())
		return
	}
	img, err := u.RandomImage()
	if err != nil {
		s.ChannelMessageSend(channel, err.Error())
		return
	}
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
