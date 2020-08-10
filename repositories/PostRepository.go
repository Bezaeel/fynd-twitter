package repositories

import "github.com/bezaeel/fynd-twitter/models"

var Posts = []models.Post{
	{
		ID:      1,
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
	// Find product
	for _, product := range Posts {
		if int(product.ID) == id {
			post = product
		}
	}

	return post, nil
}
