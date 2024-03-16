package posts

import (
	"fmt"
	"log"

	"github.com/UnikMask/gofeedsite/databases"
	"github.com/UnikMask/gofeedsite/model"
)

const (
	GET_POST_STATUS_OK           = 0
	GET_POST_STATUS_NOT_FOUND    = 1
	GET_POST_STATUS_INTERNAL_ERR = 2
)

func GetPost(post_id int, user_id int) (model.Post, int) {
	p := model.Post{}
	args := []any{&p.Id, &p.UserId, &p.Username, &p.Content, &p.PostedAt, &p.Likes, &p.Liked}

	ok, err := databases.QueryRow("databases/fetch_post.sql", []any{post_id, user_id}, args)
	if err != nil {
		log.Printf("Error fetch post %d: %v", post_id, err)
		return model.Post{}, GET_POST_STATUS_INTERNAL_ERR
	} else if !ok {
		return model.Post{}, GET_POST_STATUS_NOT_FOUND
	}
	return p, GET_POST_STATUS_OK
}

func LikePost(post model.Post, u model.UserAuth) bool {
	stmt := "databases/like_post.sql"
	if post.Liked {
		stmt = "databases/unlike_post.sql"
	}
	err := databases.ExecutePreparedStatement(stmt, u.Id, post.Id)
	if err != nil {
		fmt.Printf("Error occured toggling (file = %v) post %d as user %d: %v\n", stmt, post.Id, u.Id, err)
		return false
	}
	return true
}

func GetLikes(post_id int) (int, bool) {
	likes := 0
	found, err := databases.QueryRow("databases/get_likes.sql", []any{post_id}, []any{&likes})
	if err != nil {
		fmt.Printf("Error getting likes from post %d: %v", post_id, err)
		return 0, false
	} else if !found {
		return 0, false
	}
	return likes, true
}
