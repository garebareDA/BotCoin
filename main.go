package main

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/garebareDA/BotCoin/database"
	"github.com/garebareDA/BotCoin/handler"
)

func main() {
	db, err := database.ConnectDB()
	if err != nil {
		log.Println(err)
	}

	db.AutoMigrate(&database.User{})

	dg, err := discordgo.New("Bot" + "token")
	if err != nil {
		log.Println("error creating Discord session, ", err)
		return
	}

	dg.AddHandler(handler.AddUser)

	err = dg.Open()
	if err != nil {
		log.Println("error opening connection, ", err)
	}
}
