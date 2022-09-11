package handlers

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	bet "github.com/meximonster/go-discordbot/bet"
	cnt "github.com/meximonster/go-discordbot/content"
	"github.com/meximonster/go-discordbot/meme"
	"github.com/meximonster/go-discordbot/queries"
)

var (
	generalBetMsgConf  *betMsgSrc
	poloMsgConf        *betMsgSrc
	parolesOnlyChannel string
)

type betMsgSrc struct {
	ChannelID string
	UserID    string
}

func MessageConfigInit(generalBetAdmin string, poloBetAdmin string, generalBetChannel string, poloBetChannel string, parolesChannel string) {
	generalBetMsgConf = &betMsgSrc{
		ChannelID: generalBetChannel,
		UserID:    generalBetAdmin,
	}
	poloMsgConf = &betMsgSrc{
		ChannelID: poloBetChannel,
		UserID:    poloBetAdmin,
	}
	parolesOnlyChannel = parolesChannel
}

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself.
	if m.Author.ID == s.State.User.ID {
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

	if strings.HasPrefix(m.Content, "!add") {
		addImage(m.Content, m.ChannelID, s)
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

	if m.Content == "!users" {
		serveUsers(m.Content, m.ChannelID, s)
		return
	}

	if m.Content == "!pets" {
		servePets(m.Content, m.ChannelID, s)
		return
	}

	if m.Content == "!emotes" {
		serveEmotes(m.Content, m.ChannelID, s)
		return
	}

	if m.ChannelID == parolesOnlyChannel {
		checkForParola(m.Content, m.ChannelID, m.Attachments, s)
		return
	}

	if (m.ChannelID == generalBetMsgConf.ChannelID && m.Author.ID == generalBetMsgConf.UserID) || (m.ChannelID == poloMsgConf.ChannelID && m.Author.ID == poloMsgConf.UserID) {
		checkForBet(m.ChannelID, m.Author.ID, m.Content, s)
		return
	}

	if (m.ChannelID == generalBetMsgConf.ChannelID || m.ChannelID == poloMsgConf.ChannelID) && strings.HasPrefix(m.Content, "!bet ") {
		checkForBetQuery(m.Content, m.ChannelID, s)
		return
	}

	if (m.ChannelID == generalBetMsgConf.ChannelID || m.ChannelID == poloMsgConf.ChannelID) && strings.HasPrefix(m.Content, "!betsum ") {
		checkForBetSumQuery(m.Content, m.ChannelID, s)
		return
	}

	if strings.HasPrefix(m.Content, "!set") {
		setContent(m.Content, m.ChannelID, s)
		return
	}

	if strings.HasPrefix(m.Content, "!") {
		checkForContent(m.Content, m.ChannelID, s)
	}

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

func setContent(content string, channel string, s *discordgo.Session) {
	input := strings.Split(content, " ")
	if len(input) != 3 {
		s.ChannelMessageSend(channel, "wrong parameters")
		return
	}
	name := input[1]
	cntType := input[2]
	err := cnt.Set(name, cntType)
	if err != nil {
		s.ChannelMessageSend(channel, err.Error())
	}
}

func addImage(content string, channel string, s *discordgo.Session) {
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
	c, err := cnt.GetOne(input[1])
	if err != nil {
		s.ChannelMessageSend(channel, err.Error())
		return
	}
	err = cnt.AddImage(c, imgText, input[2])
	if err != nil {
		s.ChannelMessageSend(channel, err.Error())
	}
}

func serveGitURL(content string, channel string, s *discordgo.Session) {
	s.ChannelMessageSend(channel, "https://github.com/meximonster/go-discordbot")
}

func serveUsers(content string, channel string, s *discordgo.Session) {
	users := cnt.List("user")
	if len(users) == 0 {
		s.ChannelMessageSend(channel, "no users configured")
		return
	}
	var str string
	cnt := 0
	for _, u := range users {
		str = str + fmt.Sprintf("%d. %s\n", cnt+1, u)
		cnt++
	}
	result := "Configured users are:\n" + str
	s.ChannelMessageSend(channel, result)
}

func servePets(content string, channel string, s *discordgo.Session) {
	pets := cnt.List("pet")
	if len(pets) == 0 {
		s.ChannelMessageSend(channel, "no pets configured")
		return
	}
	var str string
	cnt := 0
	for _, p := range pets {
		str = str + fmt.Sprintf("%d. %s\n", cnt+1, p)
		cnt++
	}
	result := "Configured pets are:\n" + str
	s.ChannelMessageSend(channel, result)
}

func serveEmotes(content string, channel string, s *discordgo.Session) {
	emotes := cnt.List("emote")
	if len(emotes) == 0 {
		s.ChannelMessageSend(channel, "no emotes configured")
		return
	}
	var str string
	cnt := 0
	for _, e := range emotes {
		str = str + fmt.Sprintf("%d. %s\n", cnt+1, e)
		cnt++
	}
	result := "Configured emotes are:\n" + str
	s.ChannelMessageSend(channel, result)
}

func serveMeme(content string, channel string, s *discordgo.Session) {
	link, url, err := meme.Random()
	if err != nil {
		s.ChannelMessageSend(channel, err.Error())
		return
	}
	respondWithEmbed(channel, link, url, s)
}

func checkForParola(content string, channel string, attachments []*discordgo.MessageAttachment, s *discordgo.Session) {
	if len(attachments) > 0 || strings.HasPrefix(content, "https://www.stoiximan.gr/mybets/") {
		s.ChannelMessageSend(channel, "@everyone possible parola was just posted.")
	}
}

func checkForBet(channel string, author string, content string, s *discordgo.Session) {
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

func checkForContent(content string, channel string, s *discordgo.Session) {
	str := strings.TrimPrefix(content, "!")
	c := cnt.Get()
	if _, ok := c[str]; !ok {
		return
	}
	respondWithRandomImage(str, channel, s)
}

func checkForBetQuery(content string, channel string, s *discordgo.Session) {
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

func checkForBetSumQuery(content string, channel string, s *discordgo.Session) {
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

func respondWithRandomImage(name string, channel string, s *discordgo.Session) {
	c, err := cnt.GetOne(name)
	if err != nil {
		s.ChannelMessageSend(channel, err.Error())
		return
	}
	img, err := cnt.RandomImage(c)
	if err != nil {
		s.ChannelMessageSend(channel, err.Error())
		return
	}
	respondWithEmbed(channel, img.Text, img.Url, s)
}

func respondWithEmbed(channel string, title string, imageURL string, s *discordgo.Session) {
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
	if channel == generalBetMsgConf.ChannelID {
		table = "bets"
	} else {
		table = "polo_bets"
	}
	return table
}
