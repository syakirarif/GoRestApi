package model

import (
	"fmt"
	"strings"
)

type User struct {
	Username string `json:"username" gorm:"username"`
	Email    string `json:"email" gorm:"email"`
}

//alias name of struct
type PostItem struct {
	Title string `json:"title" gorm:"title"`
}

func (item PostItem) TableName() string {
	return "posting"
}

func (u User) ValidationUser() error {
	if len(strings.Trim(u.Username, " ")) < 1 {
		return fmt.Errorf("masukkan username")
	}

	if len(strings.Trim(u.Email, " ")) < 1 {
		return fmt.Errorf("masukkan email")
	}

	return nil
}

type Activities struct {
	Running int `json:"running" gorm:"running"`
	Walking int `json:"walking" gorm:"walking"`
}

type UserList struct {
	Data []struct {
		ID      int    `json:"id"`
		Name    string `json:"name"`
		Address string `json:"address"`
	} `json:"data"`
	Message string `json:"message"`
}

type UserJWT struct {
	UserName  string
	FirstName string
	LastName  string
}
