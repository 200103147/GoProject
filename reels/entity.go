package reels

import "time"

type Reels struct {
	Id               int
	Name_user        string
	Email_user       string
	Image_url        string

	Video_post       string
	Description_post string
	Like_post        int
	
	CreatedAt        time.Time
	UpdateAt         time.Time
}
