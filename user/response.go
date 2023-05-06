package user

type UserResponse struct {
	Id         int    `json:"id"`
	Name_user  string `json:"name_user"`
	Email_user string `json:"email_user"`
	Image_url  string `json:"image_url"`
	Following  int    `json:"following"`
	Followers  int    `json:"followers"`
	Post       int    `json:"post"`
}
