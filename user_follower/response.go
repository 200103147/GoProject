package userfollower

import "time"

type UserFollowerResponse struct {
	Id         int    `json:"id"`
	User_id    string `json:"user_id"` // link ke user
	Name_user  string `json:"name_user"`
	Email_user string `json:"email_user"`
	Image_url  string `json:"image_url"`

	CreatedAt time.Time `json:"create_at"`
	UpdateAt  time.Time `json:"update_at"`
}
