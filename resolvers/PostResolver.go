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

var AddPost = &graphql.Field{
	Type:        models.AddPostType,
	Description: "Add post",
	Args: graphql.FieldConfigArgument{
		"userId": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"message": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		userId, ok := p.Args["userId"].(int)
		mesage, messageOk := p.Args["message"].(string)

		if !ok {
			return nil, nil
		}
		if !messageOk {
			return nil, nil
		}
		return Repo.AddPost(userId, mesage)
	},
}
