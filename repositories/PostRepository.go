package repositories

import (
	"errors"
	"github.com/bezaeel/fynd-twitter/models"
)

var Posts = []models.Post{
	{
		ID:      1,
		UserId:  1,
		Message: "First Post",
		Comment: []models.Comment{
			{
				ID:      1,
				Message: "First comment",
			},
			{
				ID:      2,
				Message: "Second comment on first post",
			},
		},
	},
	{
		ID:      2,
		Message: "Second Post",
		Comment: []models.Comment{
			{
				ID:      2,
				Message: "Second comment",
			},
		},
	},
}

func GetAllPosts() []models.Post {
	return Posts
}

func GetPost(id int) (models.Post, error) {
	var post models.Post
	for _, postItem := range Posts {
		if int(postItem.ID) == id {
			post = postItem
		}
	}
	return post, nil
}

func AddPost(userId int, message string) (models.Post, error) {
	//user must exist before posting
	for _, user := range users {
		if user.Id != userId {
			return models.Post{}, errors.New("invalid userId")
		}
		break
	}
	post := models.Post{
		ID:      2,
		UserId:  userId,
		Message: message,
		Comment: nil,
	}
	var Posts = append(Posts, post)
	//log.Println(Posts)
	return Posts[2], nil
}
