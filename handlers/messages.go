package handlers

import (
	"fmt"
	"io/fs"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	bet "github.com/meximonster/go-discordbot/bet"
	"github.com/meximonster/go-discordbot/telegram"
)

var (
	parolesOnlyChannel string
)

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself.
	if m.Author.ID == s.State.User.ID {
		return
	}

	if bet.IsBetCandidate(m.Author.ID, m.ChannelID) && bet.IsBet(m.Content) {
		betNotify(m.ChannelID, m.ID, m.Author.ID, m.Content, s, m.Author.Username)
		return
	}

	if strings.HasPrefix(m.Content, "!roll") {
		rng(m.Author.Username, m.Content, m.ChannelID, s)
		return
	}

	if m.Content == "!git" {
		serveGitURL(m.Content, m.ChannelID, s)
		return
	}

	if bet.IsBetChannel(m.ChannelID) && strings.HasPrefix(m.Content, "!bet ") {
		betQuery(m.Content, m.ChannelID, s)
		return
	}

	if bet.IsBetChannel(m.ChannelID) && strings.HasPrefix(m.Content, "!betsum ") {
		betSum(m.Content, m.ChannelID, s)
		return
	}

	if bet.IsBetChannel(m.ChannelID) && m.Content == "!open" {
		serveOpenBets(m.ChannelID, s)
		return
	}

	if bet.IsBetChannel(m.ChannelID) && m.Content == "!clearbets" {
		clearBets()
		return
	}

	if bet.IsBetChannel(m.ChannelID) && strings.HasPrefix(m.Content, "!clear ") {
		clearBet(m.Content, m.ChannelID, s)
		return
	}

	if m.ChannelID == parolesOnlyChannel && (len(m.Attachments) > 0 || strings.HasPrefix(m.Content, "https://www.stoiximan.gr/mybets/")) {
		parolaNotify(m.Content, m.ChannelID, m.Attachments, s)
		return
	}

	if strings.HasPrefix(m.Content, "!") && (strings.TrimPrefix(m.Content, "!") == "users" || strings.TrimPrefix(m.Content, "!") == "pets" || strings.TrimPrefix(m.Content, "!") == "emotes") {
		listContent(m.Content, m.ChannelID, s)
		return
	}

	if strings.HasPrefix(m.Content, "!") && len(strings.Split(m.Content, " ")) == 1 {
		checkForContent(m.Content, m.ChannelID, s)
		return
	}

}

func InitChannels(ch string) {
	parolesOnlyChannel = ch
}

func listContent(content string, channel string, s *discordgo.Session) {
	cnt := strings.TrimPrefix(content, "!")
	dirs, err := os.ReadDir("./static/" + cnt)
	if err != nil {
		s.ChannelMessageSend(channel, err.Error())
	}
	var str string
	for i, dir := range dirs {
		str += fmt.Sprintf("%d. %s", i+1, dir.Name()+"\n")
	}
	_, err = s.ChannelMessageSend(channel, str)
	if err != nil {
		s.ChannelMessageSend(channel, err.Error())
	}
}

func checkForContent(content string, channel string, s *discordgo.Session) {
	name := strings.TrimPrefix(content, "!")
	imageDirs, err := os.ReadDir("./static/")
	if err != nil {
		s.ChannelMessageSend(channel, err.Error())
	}
	for _, dir := range imageDirs {
		if dir.IsDir() {
			fileDirs, err := os.ReadDir("./static/" + dir.Name() + "/")
			if err != nil {
				s.ChannelMessageSend(channel, err.Error())
			}
			for _, file := range fileDirs {
				if file.IsDir() && file.Name() == name {
					imageFiles, err := os.ReadDir("./static/" + dir.Name() + "/" + name)
					if err != nil {
						s.ChannelMessageSend(channel, err.Error())
					}
					var img fs.DirEntry
					if len(imageFiles) > 1 {
						rand.Seed(time.Now().UnixNano())
						img = imageFiles[rand.Intn(len(imageFiles)-1)+1]
					} else {
						img = imageFiles[0]
					}
					title := "**" + strings.Split(img.Name(), ".")[0] + "**"
					f, err := os.Open("./static/" + dir.Name() + "/" + name + "/" + img.Name())
					if err != nil {
						s.ChannelMessageSend(channel, err.Error())
					}
					_, err = s.ChannelMessageSendComplex(channel, &discordgo.MessageSend{
						Content: title,
						File: &discordgo.File{
							Name:   img.Name(),
							Reader: f,
						},
					})
					if err != nil {
						s.ChannelMessageSend(channel, err.Error())
					}
					break
				}
			}
		}
	}
}

func rng(username string, content string, channel string, s *discordgo.Session) {
	input := strings.Split(content, " ")
	if len(input) != 2 {
		s.ChannelMessageSend(channel, "usage: !roll <number>, result will be in range [1, <number>]")
		return
	}
	strNum := input[1]
	max, err := strconv.Atoi(strNum)
	if err != nil {
		s.ChannelMessageSend(channel, "number must be an integer")
		return
	}
	if max == 0 || max == 1 {
		s.ChannelMessageSend(channel, "no rng here")
		return
	}
	rand.Seed(time.Now().UnixNano())
	s.ChannelMessageSend(channel, fmt.Sprintf("%s rolled %d", username, rand.Intn(max-1)+1))
}

func serveGitURL(content string, channel string, s *discordgo.Session) {
	s.ChannelMessageSend(channel, "https://github.com/meximonster/go-discordbot")
}

func parolaNotify(content string, channel string, attachments []*discordgo.MessageAttachment, s *discordgo.Session) {
	s.ChannelMessageSend(channel, "@everyone possible parola was just posted.")
}

func betNotify(channel string, messageID string, author string, content string, s *discordgo.Session, username string) {
	b, err := bet.Decouple(content, "")
	if err != nil {
		s.ChannelMessageSend(channel, err.Error())
		return
	}
	bet.AddOpen(messageID, b)
	s.ChannelMessageSend(channel, fmt.Sprintf("%s %s %du @everyone", b.Team, b.Prediction, b.Size))
	go telegram.NewForwardMessage(fmt.Sprintf("%s: %s %s %du @%v", username, b.Team, b.Prediction, b.Size, b.Odds)).Forward()
}

func betQuery(content string, channel string, s *discordgo.Session) {
	table := bet.GetTableFromChannel(channel)
	q := bet.Parse(content, table)
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

func betSum(content string, channel string, s *discordgo.Session) {
	table := bet.GetTableFromChannel(channel)
	q := bet.ParseSum(content, table)
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

func serveOpenBets(channel string, s *discordgo.Session) {
	bets := bet.GetOpen()
	if len(bets) == 0 {
		s.ChannelMessageSend(channel, "no open bets")
		return
	}
	res := bet.FormatOpenBets()
	s.ChannelMessageSend(channel, res)
}

func clearBets() {
	bet.ClearAll()
}

func clearBet(content string, channel string, s *discordgo.Session) {
	input := strings.Split(content, " ")
	if len(input) != 2 {
		s.ChannelMessageSend(channel, "wrong parameters")
	}
	bet.Settle(input[1])
}
