package explore

import "time"

type Explore struct {
	Id               int
	Name_user        string
	Email_user       string
	Image_url        string

	Image_post       string
	Description_post string
	Like_post        int
	
	CreatedAt        time.Time
	UpdateAt         time.Time
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

type UserFollower struct {
	Id         int
	User_id    string // link ke user
	Name_user  string
	Email_user string
	Image_url  string

	CreatedAt time.Time
	UpdateAt  time.Time
}