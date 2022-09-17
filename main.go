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
	"github.com/meximonster/go-discordbot/bet"
	"github.com/meximonster/go-discordbot/configuration"
	"github.com/meximonster/go-discordbot/content"
	"github.com/meximonster/go-discordbot/graph"
	"github.com/meximonster/go-discordbot/handlers"
	"github.com/meximonster/go-discordbot/server"
)

var c *configuration.Config

func init() {
	err := configuration.Load()
	if err != nil {
		log.Fatal("error loading configuration: ", err)
	}

	c = configuration.Read()

	db, err := sqlx.Connect("postgres", fmt.Sprintf("postgres://127.0.0.1/postgres?sslmode=disable&user=postgres&password=%s", c.POSTGRES_PASS))
	if err != nil {
		log.Fatal("error connecting to db: ", err)
	}

	bet.NewDB(db)
	content.NewDB(db)

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

}

func main() {

	// Create a new session.
	dg, err := discordgo.New("Bot " + c.BotToken)
	if err != nil {
		log.Fatal("error creating session: ", err)
	}

	// Add handlers for message and reaction events.
	dg.AddHandler(handlers.MessageCreate)
	dg.AddHandler(handlers.ReactionCreate)

	// Add all intents that don't require privileges.
	dg.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	err = dg.Open()
	if err != nil {
		log.Fatal("error opening connection: ", err)
	}

	bet.InitAdmins(c.Admins)
	handlers.ParolesOnlyChannel = c.ParolesOnlyChannel

	for _, adm := range c.Admins {
		err := graph.Generate(adm.Table, adm.ExtraGraphs)
		if err != nil {
			log.Fatal("error generating graphs: ", err)
		}
		go graph.Schedule(adm.Table, adm.ExtraGraphs)
	}

	err = content.Load()
	if err != nil {
		log.Fatal("error loading content: ", err)
	}

	go func() {
		err := server.Run()
		if err != nil {
			log.Fatal("http server returned error: ", err)
		}
	}()

	// Create signaling for process termination.
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Gracefully stop session and close db connections and running routines.
	graph.Done()
	server.Close()
	bet.CloseDB()
	content.CloseDB()
	dg.Close()
}
