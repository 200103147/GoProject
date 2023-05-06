package reels

type ReelsRequest struct {
	Name_user        string      `json:"name_user" binding:"required"`
	Email_user       string      `json:"email_user" binding:"required"`
	Image_url        string      `json:"image_url" binding:"required"`
	
	Video_post       string      `json:"video_post" binding:"required"`
	Description_post string      `json:"description_post" binding:"required"`
	Like_post        int         `json:"like_post" binding:"required,number"`
}
