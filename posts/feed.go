package posts

import (
	"github.com/UnikMask/gofeedsite/databases"
	"github.com/UnikMask/gofeedsite/model"
)

const (
	LIMIT_POST_PER_PAGE = 10
)

func GetPosts(user_id int, page int) ([]model.Post, error) {
	rows, err := databases.Query("databases/fetch_feed.sql", user_id, LIMIT_POST_PER_PAGE*(page-1))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	posts := []model.Post{}
	for hasNext := true; hasNext; {
		p := model.Post{}
		hasNext, err = rows.ScanNext(&p.Id, &p.UserId, &p.Username, &p.Content, &p.PostedAt, &p.Likes, &p.Liked, &p.Followed)
		if hasNext {
			posts = append(posts, p)
		}
	}
	return posts, err
}
