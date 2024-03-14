package posts

import (
	"fmt"
	"log"
	"strconv"

	"github.com/UnikMask/gofeedsite/databases"
	"github.com/UnikMask/gofeedsite/model"
)

const (
	GET_POST_STATUS_OK           = 0
	GET_POST_STATUS_NOT_FOUND    = 1
	GET_POST_STATUS_INTERNAL_ERR = 2
)

func GetPost(id int) (model.Post, int) {
	p := model.Post{}
	args := []any{&p.Id, &p.UserId, &p.Username, &p.Content, &p.Likes}

	ok, err := databases.QueryRow("databases/fetch_post.sql", []any{id}, args)
	if err != nil {
		log.Printf("Error fetch post %d: %v", id, err)
		return model.Post{}, GET_POST_STATUS_INTERNAL_ERR
	} else if !ok {
		return model.Post{}, GET_POST_STATUS_NOT_FOUND
	}
	return p, GET_POST_STATUS_OK
}

func GetEndpoint(p model.Post, endpoint string) string {
    return model.ENDPOINT_POSTS + "/" + fmt.Sprintf("%d", p.Id) + endpoint;
}

func GetLikesString(p model.Post) string {
    magnitudes := []string{"", "K", "M", "B", "Q"}
    likes, i := float64(p.Likes), 0
    for likes >= 1000 {
        likes /= 1000
        i++
    }
    return fmt.Sprintf("%s%s", strconv.FormatFloat(likes, 'f', -1, 64), magnitudes[i])
}
