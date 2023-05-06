package userfollower

import "time"

type UserFollower struct {
	Id         int
	User_id    string // link ke user
	Name_user  string
	Email_user string
	Image_url  string

	CreatedAt time.Time
	UpdateAt  time.Time
}

type User struct {
	Id         int
	Name_user  string
	Email_user string
	Image_url  string
	Post       int
	Following  int
	Followers  int
}
