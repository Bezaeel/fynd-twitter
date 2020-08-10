package resolvers

import (
	"github.com/bezaeel/fynd-twitter/models"
	Repo "github.com/bezaeel/fynd-twitter/repositories"
	"github.com/graphql-go/graphql"
)

var ListPosts = &graphql.Field{
	Type:        graphql.NewList(models.PostType),
	Description: "Get posts",
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		return Repo.GetAllPosts(), nil
	},
}

var GetPost = &graphql.Field{
	Type:        models.PostType,
	Description: "Get post by id",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		id, ok := p.Args["id"].(int)
		if ok {
			return Repo.GetPost(id)
		}
		return nil, nil
	},
}

//
//import "github.com/bezaeel/fynd-twitter/models"
//
//type PostResolver struct {
//	RootResolver
//}
//
//type RootResolver struct {
//
//}
//var Posts = &[]models.Post{
//	{
//		Id: "id-0",
//		UserId: "2",
//		Message: "Ask Talabi..",
//	},
//}
//
//
//func (r *RootResolver) Posts() (*[]models.Post, error){
//	return Posts, nil
//}
//
//func
