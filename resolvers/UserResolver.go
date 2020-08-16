package resolvers

import (
	"github.com/bezaeel/fynd-twitter/models"
	Repo "github.com/bezaeel/fynd-twitter/repositories"
	"github.com/graphql-go/graphql"
	"log"
)

var ListUsers = &graphql.Field{
	Type:        graphql.NewList(models.PostType),
	Description: "Get users",
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		return Repo.GetAllUsers(), nil
	},
}

var GetUser = &graphql.Field{
	Type:        models.UserType,
	Description: "Get user by id",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		id, ok := p.Args["id"].(int)
		if ok {
			return Repo.GetUser(id)
		}
		return nil, nil
	},
}

var AddFollower = &graphql.Field{
	Type:        models.AddFollowerType,
	Description: "Add follower to a user",
	Args: graphql.FieldConfigArgument{
		"UserId": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"FollowerId": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		log.Println(p.Args)
		userId, _ := p.Args["UserId"].(int)
		followerId, followerOk := p.Args["FollowerId"].(int)
		user := models.User{}
		if followerOk {
			user, _ = Repo.AddFollower(userId, followerId)
		}
		//log.Println(user)
		return user, nil
	},
}
