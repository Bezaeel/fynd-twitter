package models

import "github.com/graph-gophers/graphql-go"

type Post struct {
	Id graphql.ID `json:"id"`
	UserId int `json:"user_id"`
	Message interface{} `json:"message"` //could be string or picture or both
	Comments []Post `json:"comment"`
}