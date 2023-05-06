package like

type LikeResponse struct {
	Email_user string `json:"email_user"`
	To_id      string `json:"to_id"`
	Post_user  string `json:"post_user"`
	Is_Like    string `json:"is_like"`
}