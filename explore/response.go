package explore

import "time"

type ExploreResponse struct {
	Id               int       `json:"id"`
	Name_user        string    `json:"name_user"`
	Email_user       string    `json:"email_user"`

	Image_post       string    `json:"image_post"`
	Description_post string    `json:"description_post"`
	Like_post        int       `json:"like_post"`
	Image_url        string    `json:"image_url"`

	CreatedAt        time.Time `json:"create_at"`
	UpdateAt         time.Time `json:"update_at"`
}
