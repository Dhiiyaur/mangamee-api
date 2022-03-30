package models

import (
	"time"
)

type User struct {
	Username    string    `json:"username" validate:"required"`
	Password    string    `json:"password" validate:"required"`
	Email       string    `json:"email"`
	CreatedDate time.Time `json:"createdDate"`
	Token       string    `json:"token"`
	LastLogin   time.Time `json:"lastLogin"`
}

type UserHistoryDelete struct {
	DeleteMangaID string `json:"MangaID" validate:"required"`
}

type UserHistory struct {
	Username      string             `json:"username"`
	UserMangaData []UserInputHistory `json:"UserMangaData"`
}

type UserInputHistory struct {
	MangaCover       string `json:"MangaCover"`
	MangaLink        string `json:"MangaLink"`
	MangaTitle       string `json:"MangaTitle"`
	MangaLastChapter string `json:"MangaLastChapter"`
	MangaLastRead    string `json:"MangaLastRead"`
	MangaSource      string `json:"MangaSource"`
}
