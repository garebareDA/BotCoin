package handler

import (
	"fmt"
	"log"
	"strconv"
	"strings"

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
	if err := db.Where("user_id = ?", m.Author.ID).First(&user).Error; err != nil {
		s.ChannelMessageSend(m.ChannelID, "口座が登録されていません\n")
		return
	}
	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("現在 %d コイン所持しています", user.Money))
}

func AddCoin(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content[0:1] != "/" {
		return
	}

	if m.Content[1:4] != "add" {
		return
	}

	sp := strings.Split(m.Content[5:], " ")
	i, err := strconv.Atoi(sp[0])
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "値が不正です")
		return
	}

	if len(m.Mentions) < 1 {
		s.ChannelMessageSend(m.ChannelID, "コインを追加する人を指定してください")
		return
	}

	db, err := database.ConnectDB()
	if err != nil {
		log.Println(err)
		s.ChannelMessageSend(m.ChannelID, "データベースエラー\n")
		return
	}

	for _, m := range m.Mentions {
		var user database.User
		db.Where("user_id = ?", m.ID).First(&user)
		user.Money += i
		log.Println(user)
		db.Save(user)
	}

	s.ChannelMessageSend(m.ChannelID, "コインの追加が完了しました")
}

func SubCoin(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content[0:1] != "/" {
		return
	}

	if m.Content[1:4] != "sub" {
		return
	}

	sp := strings.Split(m.Content[5:], " ")
	i, err := strconv.Atoi(sp[0])
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "値が不正です")
		return
	}

	if len(m.Mentions) < 1 {
		s.ChannelMessageSend(m.ChannelID, "コインを削減する人を指定してください")
		return
	}

	db, err := database.ConnectDB()
	if err != nil {
		log.Println(err)
		s.ChannelMessageSend(m.ChannelID, "データベースエラー\n")
		return
	}

	for _, m := range m.Mentions {
		var user database.User
		db.Where("user_id = ?", m.ID).First(&user)
		user.Money -= i
		log.Println(user)
		db.Save(user)
	}

	s.ChannelMessageSend(m.ChannelID, "コインの削減が完了しました")
}
