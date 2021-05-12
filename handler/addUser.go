package handler

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func AddUser(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	message := m.Content
	if message[0:1] != "/" {
		return
	}

	if message[1:] != "logup" {
		return
	}

	name := s.State.User.Username
	id := s.State.User.ID
	money := 1000

	log.Println("logup " + name + " " + id)

	//TODO DBに登録
}
