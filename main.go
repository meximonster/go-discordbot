package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/meximonster/go-discordbot/bet"
	"github.com/meximonster/go-discordbot/configuration"
	"github.com/meximonster/go-discordbot/graph"
	"github.com/meximonster/go-discordbot/handlers"
	"github.com/meximonster/go-discordbot/server"
	"github.com/meximonster/go-discordbot/telegram"
)

var (
	c  *configuration.Config
	dg *discordgo.Session
)

func init() {
	err := configuration.Load()
	if err != nil {
		log.Fatal("error loading configuration: ", err)
	}

	c = configuration.Read()

	db, err := sqlx.Connect("postgres", fmt.Sprintf("postgres://%s/postgres?sslmode=disable&user=%s&password=%s&timezone=Europe/Athens", c.POSTGRES_HOST, c.POSTGRES_USER, c.POSTGRES_PASS))
	if err != nil {
		log.Fatal("error connecting to db: ", err)
	}

	bet.NewDB(db)

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	// Create a new session.
	dg, err = discordgo.New("Bot " + c.BotToken)
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

}

func main() {

	bet.InitAdmins(c.Admins)
	handlers.InitChannels(c.ParolesOnlyChannel)

	for _, adm := range c.Admins {
		err := graph.Generate(adm.Name, adm.Table, adm.ExtraGraphs)
		if err != nil {
			log.Println("error generating graphs ", err)
		}
		go graph.Schedule(adm.Name, adm.Table, adm.ExtraGraphs)
	}

	err := bet.LoadOpen()
	if err != nil {
		log.Println("error loading open bets: ", err)
	}

	go func() {
		err := server.Run()
		if err != nil {
			log.Println("http server returned error: ", err)
		}
	}()

	telegram.NewForwardMechanism(c.FORWARD_ENDPOINT)

	log.Println("up and running!")

	// Create signaling for process termination.
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Gracefully stop session and close db connections and running routines.
	err = bet.SaveOpen()
	if err != nil {
		log.Println("error saving open bets: ", err)
	}
	graph.Done()
	server.Close()
	bet.CloseDB()
	dg.Close()
}
