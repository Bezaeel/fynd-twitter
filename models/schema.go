package models

import "github.com/graphql-go/graphql"

var CommentType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Comment",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"message": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var PostType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Post",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"message": &graphql.Field{
				Type: graphql.String,
			},
			"comment": &graphql.Field{
				Type: graphql.NewList(CommentType),
			},
		},
	},
)

var UserType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"email": &graphql.Field{
				Type: graphql.String,
			},
			"posts": &graphql.Field{
				Type: graphql.NewList(PostType),
			},
			"followers": &graphql.Field{
				Type: graphql.NewList(FollowerType),
			},
		},
	},
)

var FollowerType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Follower",
		Fields: graphql.Fields{
			"UserId": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)
