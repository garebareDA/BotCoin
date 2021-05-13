package database

import (
	"log"
	"time"
)

type User struct {
	ID       int `gorm:"AUTO_INCREMENT" gorm:"primary_key"`
	UserID   string
	UserName string
	Money    int
}

func DateUpdate() {
	time.Sleep(time.Hour * 24)
	log.Println("add money")

	db, err := ConnectDB()
	if err != nil {
		log.Println("database connect error")
		return
	}

	var lastUser User
	db.Last(&lastUser)
	lastID := lastUser.ID

	for i := 0; lastID < i; i++ {

	}
}
