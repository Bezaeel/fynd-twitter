package models

import "github.com/graph-gophers/graphql-go"

//users
type User struct{
	Id graphql.ID `json:"id"`
	Name string `json:"name"`
	Posts []Post `json:"posts"`
	Followers []Follower `json:"followers"`
}

type Follower struct {
	UserId graphql.ID `json:"user_id"`
	//Email string `json:"email"`
}
