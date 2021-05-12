package main

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func main() {
	dg, err := discordgo.New("Bot" + "token")
	if err != nil {
		log.Println("error creating Discord session, ", err)
		return
	}

	err = dg.Open()
	if err != nil {
		log.Println("error opening connection, ", err)
	}
}
