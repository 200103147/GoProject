package comment

import "time"

type CommentRequest struct {
	Name_user  string `json:"name_user" binding:"required"`
	ToId       string `json:"to_id" binding:"required"`
	Email_user string `json:"email_user" binding:"required"`
	Image_url  string `json:"image_url" binding:"required"`

	Comment    string `json:"comment" binding:"required"`
	PostUser   string `json:"post_user" binding:"required"`
	ToUser     string `json:"to_user" binding:"required"`

	CreatedAt time.Time `json:"create_at"`
	UpdateAt  time.Time `json:"update_at"`
}
