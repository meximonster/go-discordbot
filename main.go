package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/meximonster/go-discordbot/configuration"
	"github.com/meximonster/go-discordbot/handlers"
)

func init() {
	err := configuration.Load()
	if err != nil {
		log.Fatal("error loading configuration: ", err)
	}
}

func main() {
	c := configuration.Read()

	// Create a new session.
	dg, err := discordgo.New("Bot " + c.BotToken)
	if err != nil {
		log.Fatal("error creating session: ", err)
	}

	handlers.MessageConfigInit(c.ChannelID, c.UserID)

	// Add handler for message and reaction events.
	dg.AddHandler(handlers.MessageCreate)
	dg.AddHandler(handlers.ReactionCreate)

	// Only care about receiving message events.
	dg.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	err = dg.Open()
	if err != nil {
		log.Fatal("error opening connection: ", err)
	}

	fmt.Println("Bot is running!")

	// Create signaling for process termination.
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Gracefully stop session.
	dg.Close()
}
