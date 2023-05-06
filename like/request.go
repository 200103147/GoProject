package like

type LikeRequest struct {
	Email_user string `json:"email_user" binding:"required"`
	To_id      string `json:"to_id" binding:"required"`
	Post_user  string `json:"post_user" binding:"required"`
	Is_Like    string `json:"is_like" binding:"required"`
}
