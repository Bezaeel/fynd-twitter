package models

//users
type User struct {
	Id        int64      `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Posts     []Post     `json:"posts"`
	Followers []Follower `json:"followers"`
}

type Follower struct {
	UserId int64 `json:"user_id"`
}
