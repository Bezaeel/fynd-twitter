package models

type Post struct {
	ID      int       `json:"id"`
	UserId  int       `json:"user_id"`
	Message string    `json:"message"`
	Comment []Comment `json:"comment"`
}

type Comment struct {
	ID      int    `json:"id"`
	PostId  int    `json:"post_id"`
	UserId  string `json:"user_id"`
	Message string `json:"message"`
}
