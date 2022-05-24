package handlers

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	bet "github.com/meximonster/go-discordbot/bet"
	"github.com/meximonster/go-discordbot/configuration"
	cnt "github.com/meximonster/go-discordbot/content"
	"github.com/meximonster/go-discordbot/meme"
	"github.com/meximonster/go-discordbot/queries"
)

var (
	betMsgConf      *betMsgSrc
	poloMsgConf     *betMsgSrc
	parolaChannelID string
)

type betMsgSrc struct {
	ChannelID string
	UserID    string
}

func MessageConfigInit(content []configuration.CntConfig, parolaChannel string) {
	parolaChannelID = parolaChannel
	for _, c := range content {
		switch strings.ToLower(c.Name) {
		case "pad":
			betMsgConf = &betMsgSrc{
				UserID:    c.UserID,
				ChannelID: c.ChannelID,
			}
		case "fyk":
			poloMsgConf = &betMsgSrc{
				UserID:    c.UserID,
				ChannelID: c.ChannelID,
			}
		}
	}
}

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself.
	if m.Author.ID == s.State.User.ID {
		return
	}

	rng(m.Content, m.ChannelID, s)
	setContent(m.Content, m.ChannelID, s)
	addImage(m.Content, m.ChannelID, s)
	serveGitURL(m.Content, m.ChannelID, s)
	serveMeme(m.Content, m.ChannelID, s)
	serveUsers(m.Content, m.ChannelID, s)
	servePets(m.Content, m.ChannelID, s)
	serveArtists(m.Content, m.ChannelID, s)
	checkForContent(m.Content, m.ChannelID, s)
	checkForParola(m.Content, m.ChannelID, m.Attachments, s)
	checkForBet(m.ChannelID, m.Author.ID, m.Content, s)
	checkForBetQuery(m.Content, m.ChannelID, s)
	checkForBetSumQuery(m.Content, m.ChannelID, s)
}

func rng(content string, channel string, s *discordgo.Session) {
	if strings.HasPrefix(content, "!roll") {
		input := strings.Split(content, " ")
		if len(input) != 2 {
			s.ChannelMessageSend(channel, "usage: !roll <number>, result will be in range [1, <number>]")
			return
		}
		strNum := input[1]
		max, err := strconv.Atoi(strNum)
		if err != nil {
			s.ChannelMessageSend(channel, "number must be integer")
			return
		}
		rand.Seed(time.Now().UnixNano())
		s.ChannelMessageSend(channel, fmt.Sprintf("%d", rand.Intn(max)+1))
	}
}

func setContent(content string, channel string, s *discordgo.Session) {
	if strings.HasPrefix(content, "!set") {
		input := strings.Split(content, " ")
		if len(input) != 3 {
			s.ChannelMessageSend(channel, "wrong parameters")
			return
		}
		name := input[1]
		cntType := input[2]
		if cntType == "human" {
			users := cnt.GetUsers()
			if _, ok := users[name]; ok {
				s.ChannelMessageSend(channel, fmt.Sprintf("user %s already exists", name))
				return
			}
		} else if cntType == "pet" {
			pets := cnt.GetPets()
			if _, ok := pets[name]; ok {
				s.ChannelMessageSend(channel, fmt.Sprintf("pet %s already exists", name))
				return
			}
		} else if cntType == "artist" {
			artists := cnt.GetArtists()
			if _, ok := artists[name]; ok {
				s.ChannelMessageSend(channel, fmt.Sprintf("artist %s already exists", name))
				return
			}
		} else {
			s.ChannelMessageSend(channel, "content type should be either human or pet")
			return
		}
		err := cnt.Set(name, cntType)
		if err != nil {
			s.ChannelMessageSend(channel, err.Error())
		}
	}
}

func addImage(content string, channel string, s *discordgo.Session) {
	if strings.HasPrefix(content, "!add") {
		text := strings.Split(content, "'")
		if len(text) != 3 {
			s.ChannelMessageSend(channel, "wrong parameters")
			return
		}
		imgText := text[1]
		replace := " '" + imgText + "'"
		str := strings.Replace(content, replace, "", 1)
		input := strings.Split(str, " ")
		if len(input) < 3 {
			s.ChannelMessageSend(channel, "not enough parameters")
			return
		}
		if len(input) > 3 {
			s.ChannelMessageSend(channel, "too many parameters")
			return
		}
		err := cnt.AddImage(input[1], imgText, input[2])
		if err != nil {
			s.ChannelMessageSend(channel, err.Error())
		}
	}
}

func serveGitURL(content string, channel string, s *discordgo.Session) {
	if content == "!git" {
		s.ChannelMessageSend(channel, "https://github.com/meximonster/go-discordbot")
	}
}

func serveUsers(content string, channel string, s *discordgo.Session) {
	if content == "!users" {
		users := cnt.GetUsers()
		if len(users) == 0 {
			s.ChannelMessageSend(channel, "no users configured")
			return
		}
		var str string
		cnt := 0
		for _, u := range users {
			str = str + fmt.Sprintf("%d. %s\n", cnt+1, u.Name)
			cnt++
		}
		result := "Configured users are:\n" + str
		s.ChannelMessageSend(channel, result)
	}
}

func servePets(content string, channel string, s *discordgo.Session) {
	if content == "!pets" {
		pets := cnt.GetPets()
		if len(pets) == 0 {
			s.ChannelMessageSend(channel, "no pets configured")
			return
		}
		var str string
		cnt := 0
		for _, u := range pets {
			str = str + fmt.Sprintf("%d. %s\n", cnt+1, u.Name)
			cnt++
		}
		result := "Configured pets are:\n" + str
		s.ChannelMessageSend(channel, result)
	}
}

func serveArtists(content string, channel string, s *discordgo.Session) {
	if content == "!artists" {
		artists := cnt.GetArtists()
		if len(artists) == 0 {
			s.ChannelMessageSend(channel, "no artists configured")
			return
		}
		var str string
		cnt := 0
		for _, a := range artists {
			str = str + fmt.Sprintf("%d. %s\n", cnt+1, a.Name)
			cnt++
		}
		result := "Configured artists are:\n" + str
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
	if (channel == betMsgConf.ChannelID && author == betMsgConf.UserID) || (channel == poloMsgConf.ChannelID && author == poloMsgConf.UserID) {
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

func checkForContent(content string, channel string, s *discordgo.Session) {
	if strings.HasPrefix(content, "!") {
		str := strings.TrimPrefix(content, "!")
		c := cnt.Get()
		if _, ok := c[str]; ok {
			respondWithRandomImage(str, channel, s)
		}
	}
}

func checkForBetQuery(content string, channel string, s *discordgo.Session) {
	if (channel == betMsgConf.ChannelID || channel == poloMsgConf.ChannelID) && strings.HasPrefix(content, "!bet ") {
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
	if (channel == betMsgConf.ChannelID || channel == poloMsgConf.ChannelID) && strings.HasPrefix(content, "!betsum ") {
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
	u, err := cnt.GetByName(name)
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
	if channel == betMsgConf.ChannelID {
		table = "bets"
	} else {
		table = "polo_bets"
	}
	return table
}
