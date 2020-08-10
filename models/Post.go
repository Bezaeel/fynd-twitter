package models

type Post struct {
	ID      int64     `json:"id"`
	Message string    `json:"message"`
	Comment []Comment `json:"comment"`
}

type Comment struct {
	ID      int64  `json:"id"`
	Message string `json:"message"`
}
