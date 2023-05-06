package explore

type Service interface {
	Create(exploreRequest ExploreRequest) (Explore, error)
	FindByNotUserAndOrderByLike(user string) ([]Explore, error)
	GetByEmailAndOrderByCreateAt(email string) ([]Explore, error)
	GetByUserFollowing(email_user string) ([]Explore, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Create(exploreRequest ExploreRequest) (Explore, error) {
	book := Explore {
		Name_user: exploreRequest.Name_user,
		Email_user: exploreRequest.Email_user,
		Image_url: exploreRequest.Image_url,

		Image_post: exploreRequest.Image_post,
		Description_post: exploreRequest.Description_post,
		Like_post: exploreRequest.Like_post,
	}

	newBook, err := s.repository.Create(book)

	return newBook, err
	// return s.repository.FindAll()
}

func (s *service) FindByNotUserAndOrderByLike(user string) ([]Explore, error) {
	books, err := s.repository.FindByNotUserAndOrderByLike(user)
	return books, err
}

func (s *service) GetByEmailAndOrderByCreateAt(email string) ([]Explore, error) {
	books, err := s.repository.GetByEmailAndOrderByCreateAt(email)
	return books, err
}

func (s *service) GetByUserFollowing(email string) ([]Explore, error) {
	books, err := s.repository.GetByUserFollowing(email)
	return books, err
}