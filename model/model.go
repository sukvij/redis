package model

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        int64
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

type User struct {
	BaseModel
	FirstName string `json:"first_name" `
	LastName  string `json:"last_name"`
	Gender    string `json:"gender"`
	Age       int    `json:"age"`
	Contact   string `json:"contact"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	IsMember  bool   `json:"is_member"`
	Priority  int    `json:"priority"`
	Verified  bool   `json:"verified"`
	Image     string `json:"image"`
}

func (user *User) BeforeCreate(db *gorm.DB) error {
	fmt.Println("before")
	return nil
}

func (user *User) AfterCreate(db *gorm.DB) error {
	fmt.Println("after")
	return nil
}
