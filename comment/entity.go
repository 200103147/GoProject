package comment

import "time"

type Comment struct {
	Id         int
	ToId	   string
	Name_user  string
	Email_user string
	Image_url  string
	
	Comment    string
	PostUser   string
	ToUser     string

	CreatedAt time.Time
	UpdateAt  time.Time
}
