package model

import (
	"fmt"
	"strconv"
	"time"
)

type Post struct {
	Id       int
	UserId   int
	Username string
	Content  string
	PostedAt time.Time
	Likes    int
	Liked    bool
}

func (p Post) GetDatePostedString() string {
	t := time.Now().Sub(p.PostedAt)
	if t.Seconds() < 60 {
		return fmt.Sprintf("%.0f", t.Seconds()) + "s"
	} else if t.Minutes() < 60 {
		return fmt.Sprintf("%.0f", t.Minutes()) + "m"
	} else if t.Hours() < 24 {
		return fmt.Sprintf("%.0f", t.Hours()) + "h"
	} else if t.Hours() < 24*7 {
		return fmt.Sprintf("%.0f", t.Hours()/7) + "d"
	} else if p.PostedAt.Year() == time.Now().Year() {
		return p.PostedAt.Format("02/01")
	}
	return p.PostedAt.Format("02/01/2006")
}

func GetLikesString(amount int) string {
	magnitudes := []string{"", "K", "M", "B", "Q"}
	likes, i := float64(amount), 0
	for likes >= 1000 {
		likes /= 1000
		i++
	}
	return fmt.Sprintf("%s%s", strconv.FormatFloat(likes, 'f', -1, 64), magnitudes[i])
}

func (p Post) GetEndpoint(endpoint string) string {
	return ENDPOINT_POSTS + "/" + fmt.Sprintf("%d", p.Id) + endpoint
}
