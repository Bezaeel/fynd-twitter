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
				UserId: 1,
			},
		},
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

//
//import (
//	"github.com/bezaeel/fynd-twitter/models"
//)
//
//var Posts = []models.Post{
//	{
//		Id: "id-0",
//		UserId: "2",
//		Message: "Ask Talabi..",
//	},
//}
//
//var Users = []models.User{
//	{
//		Id: "3",
//		Name: "Rotimi",
//		Email: "r@mail.co",
//		Posts: Posts,
//	},
//}
//
//type UserResponse struct {
//
//	response []models.User
//}
//
//
////all users
//func AllUsers() *UserResponse{
//	return &UserResponse{response: Users}
//}
//
////all user's post
////func GetUserPosts(userId graphql.ID){
////	for _, user := range Users{
////		if userId == user.Id{
////			b, _ := json.Marshal(user.Posts)
////			fmt.Println(string(b))
////		}
////	}
////}
//
////all user's followers
////func AllUserFollowers(userId graphql.ID) {
////	for _, user := range Users{
////		if userId == user.Id{
////			b, _ := json.Marshal(user.Followers)
////			fmt.Println(string(b))
////		}
////	}
////}
//
////all user's following
////func GetFollowing(userId graphql.ID) {
////	for _, user := range Users{
////		for _, following := range user.Followers{
////			if userId == following.UserId{
////				b, _ := json.Marshal(user)
////				fmt.Println(string(b))
////			}
////		}
////	}
////}
