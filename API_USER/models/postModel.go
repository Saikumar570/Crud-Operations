package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Name       string
	Employeeid int
	Email      string
	Comments   []Comment `gorm:"foreignKey:PostID"`
}

type Comment struct {
	gorm.Model
	Content string
	PostID  uint
}
