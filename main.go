package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/garebareDA/BotCoin/database"
	"github.com/garebareDA/BotCoin/handler"
)

var (
	Token string
)

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {
	db, err := database.ConnectDB()
	if err != nil {
		log.Println("database coonect error, ", err)
	}

	db.AutoMigrate(&database.User{})

	go database.DateUpdate()

	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		log.Println("error creating Discord session, ", err)
		return
	}

	dg.AddHandler(handler.AddUser)
	dg.AddHandler(handler.CoinCheck)
	dg.AddHandler(handler.AddCoin)
	dg.AddHandler(handler.SubCoin)

	err = dg.Open()
	if err != nil {
		log.Println("error opening connection, ", err)
		return
	}

	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGABRT, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
}
