package handlers

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	bet "github.com/meximonster/go-discordbot/bet"
	"github.com/meximonster/go-discordbot/integrations/pubg"
	"github.com/meximonster/go-discordbot/integrations/wow"
	"github.com/meximonster/go-discordbot/meme"
)

var (
	parolesOnlyChannel string
	r8mypl8Channel     string
)

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself.
	if m.Author.ID == s.State.User.ID {
		return
	}

	if bet.IsBetCandidate(m.Author.ID, m.ChannelID) && bet.IsBet(m.Content) {
		betNotify(m.ChannelID, m.ID, m.Author.ID, m.Content, s)
		return
	}

	if strings.HasPrefix(m.Content, "!last") {
		getPubgLastMatchInfo(m.Content, m.ChannelID, s)
		return
	}

	if strings.HasPrefix(m.Content, "!season") {
		getPubgSeasonStats(m.Content, m.ChannelID, s)
		return
	}

	if strings.HasPrefix(m.Content, "!ranked") {
		getPubgRankedStats(m.Content, m.ChannelID, s)
		return
	}

	if strings.HasPrefix(m.Content, "!rating") {
		getRating(m.Content, m.ChannelID, s)
		return
	}

	if strings.HasPrefix(m.Content, "!tts") {
		tts(m.Content, m.ChannelID, s)
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

	if m.Content == "!meme" {
		serveMeme(m.Content, m.ChannelID, s)
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

	if m.ChannelID == r8mypl8Channel && (len(m.Attachments)) > 0 {
		ratePlate(m.Content, m.ChannelID, m.Author.Username, s)
		return
	}

}

func InitChannels(paroles string, r8mypl8 string) {
	parolesOnlyChannel = paroles
	r8mypl8Channel = r8mypl8
}

func getPubgSeasonStats(content string, channel string, s *discordgo.Session) {
	input := strings.Split(content, " ")
	if len(input) != 3 {
		s.ChannelMessageSend(channel, "wrong parameters - usage: !stats <name> <solo/duo/squad>")
		return
	}
	stats, err := pubg.GetSeasonStats(input[1], input[2])
	if err != nil {
		s.ChannelMessageSend(channel, err.Error())
		return
	}
	s.ChannelMessageSend(channel, stats)
}

func getPubgLastMatchInfo(content string, channel string, s *discordgo.Session) {
	input := strings.Split(content, " ")
	if len(input) != 2 {
		s.ChannelMessageSend(channel, "wrong parameters - usage: !last <name>")
		return
	}
	info, err := pubg.GetLastMatchInfo(input[1])
	if err != nil {
		s.ChannelMessageSend(channel, err.Error())
		return
	}
	s.ChannelMessageSend(channel, info)
}

func getPubgRankedStats(content string, channel string, s *discordgo.Session) {
	input := strings.Split(content, " ")
	if len(input) != 3 {
		s.ChannelMessageSend(channel, "wrong parameters - usage: !stats <name> <solo/duo/squad>")
		return
	}
	stats, err := pubg.GetRankedSeasonStats(input[1], input[2])
	if err != nil {
		s.ChannelMessageSend(channel, err.Error())
		return
	}
	s.ChannelMessageSend(channel, stats)
}

func getRating(content string, channel string, s *discordgo.Session) {
	input := strings.Split(content, " ")
	if len(input) != 3 {
		s.ChannelMessageSend(channel, "wrong parameters - usage: !rating <name> <realm>")
		return
	}
	profile, err := wow.GetProfile(input[2], input[1])
	if err != nil {
		s.ChannelMessageSend(channel, err.Error())
		return
	}
	s.ChannelMessageSendEmbed(channel, &discordgo.MessageEmbed{
		Footer: &discordgo.MessageEmbedFooter{
			Text: profile,
		},
	})
}

func ratePlate(content string, channel string, username string, s *discordgo.Session) {
	rand.Seed(time.Now().UnixNano())
	s.ChannelMessageSend(channel, fmt.Sprintf("%s m8, i r8 your pl8 %d/8", username, rand.Intn(8)+1))
}

func tts(content string, channel string, s *discordgo.Session) {
	tts := strings.Replace(content, "!tts ", "", 1)
	s.ChannelMessageSendComplex(channel, &discordgo.MessageSend{
		Content: tts,
		TTS:     true,
	})
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
	s.ChannelMessageSend(channel, fmt.Sprintf("%s rolled %d", username, rand.Intn(max)+1))
}

func serveGitURL(content string, channel string, s *discordgo.Session) {
	s.ChannelMessageSend(channel, "https://github.com/meximonster/go-discordbot")
}

func serveMeme(content string, channel string, s *discordgo.Session) {
	link, url, err := meme.Random()
	if err != nil {
		s.ChannelMessageSend(channel, err.Error())
		return
	}
	respondWithEmbed(channel, link, url, s)
}

func parolaNotify(content string, channel string, attachments []*discordgo.MessageAttachment, s *discordgo.Session) {
	s.ChannelMessageSend(channel, "@everyone possible parola was just posted.")
}

func betNotify(channel string, messageID string, author string, content string, s *discordgo.Session) {
	b, err := bet.Decouple(content, "")
	if err != nil {
		s.ChannelMessageSend(channel, err.Error())
		return
	}
	bet.AddOpen(messageID, b)
	s.ChannelMessageSend(channel, fmt.Sprintf("%s %s %du @everyone", b.Team, b.Prediction, b.Size))
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

func respondWithEmbed(channel string, title string, imageURL string, s *discordgo.Session) {
	_, err := s.ChannelMessageSendEmbed(channel, &discordgo.MessageEmbed{
		Title: title,
		Image: &discordgo.MessageEmbedImage{
			URL: imageURL,
		},
	})
	if err != nil {
		s.ChannelMessageSend(channel, err.Error())
	}
}
