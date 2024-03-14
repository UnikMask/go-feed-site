package model

import "time"

type Post struct {
	Id       int
	UserId   int
	Username string
	Content  string
	PostedAt time.Time
	Likes    int
}
