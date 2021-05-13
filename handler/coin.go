package handler

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/garebareDA/BotCoin/database"
)

func CoinCheck(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	message := m.Content
	if message[0:1] != "/" {
		return
	}

	if message[1:] != "check" {
		return
	}

	db, err := database.ConnectDB()
	if err != nil {
		log.Println(err)
		s.ChannelMessageSend(m.ChannelID, "データベースエラー\n")
		return
	}

	user := database.User{}
	if err := db.First(&user).Error; err != nil {
		s.ChannelMessageSend(m.ChannelID, "口座が登録されていません\n")
		return
	}

	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("現在 %d コイン所持しています", user.Money))
}
