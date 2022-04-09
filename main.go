package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/bwmarrin/discordgo"
	"github.com/meximonster/go-discordbot/bets"
	"github.com/meximonster/go-discordbot/configuration"
	"github.com/meximonster/go-discordbot/handlers"
)

func init() {
	err := configuration.Load()
	if err != nil {
		log.Fatal("error loading configuration: ", err)
	}

	db, err := sqlx.Connect("postgres", "postgres://127.0.0.1/postgres?sslmode=disable&user=postgres&password=postgres")
	if err != nil {
		log.Fatal("error connecting to db: ", err)
	}
	bets.NewDB(db)
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)
}

func main() {
	c := configuration.Read()

	// Create a new session.
	dg, err := discordgo.New("Bot " + c.BotToken)
	if err != nil {
		log.Fatal("error creating session: ", err)
	}

	handlers.MessageConfigInit(c.ChannelID, c.UserID)

	// Add handlers for message and reaction events.
	dg.AddHandler(handlers.MessageCreate)
	dg.AddHandler(handlers.ReactionCreate)

	// Add all intents that don't require privileges.
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

	// Gracefully stop session and close db connection.
	bets.CloseDB()
	dg.Close()
}
