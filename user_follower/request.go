package userfollower

type UserFollowerRequest struct {
	User_id    string `json:"user_id" binding:"required"`
	Name_user  string `json:"name_user" binding:"required"`
	Email_user string `json:"email_user" binding:"required"`
	Image_url  string `json:"image_url" binding:"required"`
}