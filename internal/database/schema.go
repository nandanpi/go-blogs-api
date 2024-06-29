package database

import (
	"log"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	ID       int        `json:"id" gorm:"primaryKey;autoIncrement;not null;uniqueIndex"`
	Username string     `json:"username" gorm:"unique"`
	Password string     `json:"password"`
	Image    string     `json:"image"`
	Blogs    []BlogPost `json:"blogs" gorm:"foreignKey:UserID"`
	Likes    []Likes    `json:"likes" gorm:"foreignKey:UserID"`
	Comments []Comments `json:"comments" gorm:"foreignKey:UserID"`
}

type BlogPost struct {
	gorm.Model
	ID       int        `json:"id" gorm:"primaryKey;autoIncrement;not null;uniqueIndex"`
	Title    string     `json:"title"`
	Content  string     `json:"content"`
	UserID   int        `json:"user_id"`
	User     Users      `json:"user" gorm:"foreignKey:UserID"`
	Likes    []Likes    `json:"likes" gorm:"foreignKey:PostID"`
	Comments []Comments `json:"comments" gorm:"foreignKey:PostID"`
}

type Likes struct {
	gorm.Model
	ID     int      `json:"id" gorm:"primaryKey;autoIncrement;not null;uniqueIndex"`
	UserID int      `json:"user_id"`
	PostID int      `json:"post_id"`
	Post   BlogPost `json:"post" gorm:"foreignKey:PostID"`
	User   Users    `json:"user" gorm:"foreignKey:UserID"`
}

type Comments struct {
	gorm.Model
	ID      int      `json:"id" gorm:"primaryKey;autoIncrement;not null;uniqueIndex"`
	Content string   `json:"content"`
	UserID  int      `json:"user_id"`
	PostID  int      `json:"post_id"`
	User    Users    `json:"user" gorm:"foreignKey:UserID"`
	Post    BlogPost `json:"post" gorm:"foreignKey:PostID"`
}

func PushSchema(db *gorm.DB) {
	err := db.AutoMigrate(&Users{}, &BlogPost{}, &Likes{}, &Comments{})
	if err != nil {
		panic("Failed to migrate database schema")
	}
	log.Println("Migrated DB schema")
}
