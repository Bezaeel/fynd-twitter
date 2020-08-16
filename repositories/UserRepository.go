package repositories

import (
	"github.com/bezaeel/fynd-twitter/models"
)

var users = []models.User{
	{
		Id:    1,
		Name:  "Talabi",
		Email: "talabi@mail.com",
		Posts: Posts,
		Followers: []models.Follower{
			{
				UserId: 2,
			},
		},
	},
	{
		Id:        2,
		Name:      "Opemipo",
		Email:     "Opemipo@mail.com",
		Posts:     []models.Post{},
		Followers: []models.Follower{},
	},
	{
		Id:        3,
		Name:      "TalOpe",
		Email:     "T@mail.com",
		Posts:     []models.Post{},
		Followers: []models.Follower{},
	},
}

func GetAllUsers() []models.User {
	return users
}

func GetUser(id int) (models.User, error) {
	var user models.User
	// Find product
	for _, record := range users {
		if int(record.Id) == id {
			user = record
		}
	}

	return user, nil
}

func AddFollower(userId int, followerId int) (models.User, error) {
	user := models.User{}
	follower := models.Follower{UserId: followerId}
	for _, record := range users {
		if int(record.Id) == userId {
			record.Followers = append(record.Followers, follower)
			user = record
		}
	}
	return user, nil
}
