package database

import (
	"log"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	ID       uint       `json:"id" gorm:"primaryKey;autoIncrement;not null;uniqueIndex"`
	Username string     `json:"username" gorm:"unique;not null"`
	Password string     `json:"password" gorm:"not null"`
	Image    string     `json:"image"`
	Blogs    []BlogPost `json:"blogs" gorm:"foreignKey:UserID"`
	Likes    []Likes    `json:"likes" gorm:"foreignKey:UserID"`
	Comments []Comments `json:"comments" gorm:"foreignKey:UserID"`
}

type BlogPost struct {
	gorm.Model
	ID       uint       `json:"id" gorm:"primaryKey;autoIncrement;not null;uniqueIndex"`
	Title    string     `json:"title" gorm:"not null"`
	Content  string     `json:"content" gorm:"type:text;not null"`
	UserID   uint       `json:"user_id"`
	User     Users      `json:"user" gorm:"foreignKey:UserID"`
	Likes    []Likes    `json:"likes" gorm:"foreignKey:PostID"`
	Comments []Comments `json:"comments" gorm:"foreignKey:PostID"`
}

type Likes struct {
	gorm.Model
	ID     uint `json:"id" gorm:"primaryKey;autoIncrement;not null;uniqueIndex"`
	UserID uint `json:"user_id"`
	PostID uint `json:"post_id"`
}

type Comments struct {
	gorm.Model
	ID      uint   `json:"id" gorm:"primaryKey;autoIncrement;not null;uniqueIndex"`
	Content string `json:"content" gorm:"type:text;not null"`
	UserID  uint   `json:"user_id"`
	PostID  uint   `json:"post_id"`
}

func PushSchema(db *gorm.DB) {
	err := db.AutoMigrate(&Users{}, &BlogPost{}, &Likes{}, &Comments{})
	if err != nil {
		panic("Failed to migrate database schema")
	}
	log.Println("Migrated DB schema")
}
