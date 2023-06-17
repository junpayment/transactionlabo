package domain

type User struct {
	ID   int64  `gorm:"id" json:"id"`
	Name string `gorm:"name" json:"name"`
}

type Role struct {
	ID     int64  `gorm:"id" json:"id"`
	Name   string `gorm:"name" json:"name"`
	UserID int64  `gorm:"user_id" json:"user_id"`
}
