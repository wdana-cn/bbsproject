package model

import (
	"time"
)

type User struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	Name      string     `gorm:"type:varchar(20);not null"`
	Telephone string     `gorm:"varchar(11);not null;unique"`
	Password  string     `gorm:"size:255;not null"`
	Email     string     `gorm:"type:varchar(30)"`
	Avator    string     `gorm:"type:varchar(100)"`
}

type Postdata struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	Title     string     `gorm:"type:varchar(100);not null"`
	Postvalue string     `gorm:"type:varchar(20);not null"`
	Content   string     `gorm:"type:text;not null"`
	User      User       `gorm:"foreignkey:UserId;associationforeignkey:ID"`
	UserId    uint       `gorm:"not null"`
}
type Comment struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	Postid    uint       `gorm:"not null"`
	Userid    uint		 `gorm:"not null"`
	Comment   string `gorm:"type:varchar(500);"`
}
