package database

type User struct {
	ID       int `gorm:"AUTO_INCREMENT" gorm:"primary_key"`
	UserID   string
	UserName string
	Money    int
}
