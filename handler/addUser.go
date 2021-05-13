package handler

import (
	"log"

	"github.com/bwmarrin/discordgo"

	"github.com/garebareDA/BotCoin/database"
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
	db, err := database.ConnectDB()
	if err != nil {
		log.Println(err)
		s.ChannelMessageSend(m.ChannelID, "データベースエラー\n")
		return
	}

	if err := db.Where("user_id = ?", id).First(&database.User{}).Error; err != nil {
		s.ChannelMessageSend(m.ChannelID, "口座がすでに存在しています\n")
		return
	}

	user := database.User{
		UserID:   id,
		UserName: name,
		Money:    money,
	}
	db.Where(database.User{UserID: id}).FirstOrInit(&user)

	s.ChannelMessageSend(m.ChannelID, "口座の開設が完了しました\n")
}
