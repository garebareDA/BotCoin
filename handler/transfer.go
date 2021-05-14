package handler

import (
	"log"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/garebareDA/BotCoin/database"
)

func Transfer(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	message := m.Content
	if message[0:1] != "/" {
		return
	}

	if message[1:9] != "transfer" {
		return
	}

	sp := strings.Split(m.Content[9:], " ")
	i, err := strconv.Atoi(sp[0])
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "値が不正です")
		return
	}

	if len(m.Mentions) != 1 {
		s.ChannelMessageSend(m.ChannelID, "振り込み先が不正です")
		return
	}

	db, err := database.ConnectDB()
	if err != nil {
		log.Println(err)
		s.ChannelMessageSend(m.ChannelID, "データベースエラー\n")
		return
	}

	var user database.User
	db.Where("user_id = ?", m.Author.ID).First(&user)
	user.Money -= i
	db.Save(user)

	var transfer database.User
	db.Where("user_id = ?", m.Author.ID).First(&transfer)
	transfer.Money += i
	db.Save(transfer)

	s.ChannelMessageSend(m.ChannelID, "振り込みが完了しました\n")
}
