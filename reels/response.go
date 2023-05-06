package reels

import "time"

type ReelsResponse struct {
	Id               int       `json:"id"`
	Name_user        string    `json:"name_user"`
	Email_user       string    `json:"email_user"`
	Image_url        string    `json:"image_url"`

	Video_post       string    `json:"video_post"`
	Description_post string    `json:"description_post"`
	Like_post        int       `json:"like_post"`

	CreatedAt        time.Time `json:"create_at"`
	UpdateAt         time.Time `json:"update_at"`
}
