package reels

type Service interface {
	Create(reelsRequest ReelsRequest) (Reels, error)
	FindByNotUserAndOrderByLike(user string) ([]Reels, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Create(exploreRequest ReelsRequest) (Reels, error) {
	book := Reels {
		Name_user: exploreRequest.Name_user,
		Email_user: exploreRequest.Email_user,
		Image_url: exploreRequest.Image_url,

		Video_post: exploreRequest.Video_post,
		Description_post: exploreRequest.Description_post,
		Like_post: exploreRequest.Like_post,
	}

	newBook, err := s.repository.Create(book)

	return newBook, err
	// return s.repository.FindAll()
}

func (s *service) FindByNotUserAndOrderByLike(user string) ([]Reels, error) {
	books, err := s.repository.FindByNotUserAndOrderByLike(user)
	return books, err
}