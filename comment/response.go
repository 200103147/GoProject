package comment

import "time"

type CommentResponse struct {
	Id         int    `json:"id"`
	ToId       string `json:"to_id"`
	Name_user  string `json:"name_user"`
	Email_user string `json:"email_user"`
	Image_url  string `json:"image_url"`

	Comment    string `json:"comment"`
	PostUser   string `json:"post_user"`
	ToUser     string `json:"to_user"`

	CreatedAt time.Time `json:"create_at"`
	UpdateAt  time.Time `json:"update_at"`
}
