package like

type Service interface {
	Create(likeRequest LikeRequest) (Like, error)
	FindByUserLike(id string, email_user string) ([]Like, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Create(likeRequest LikeRequest) (Like, error) {
	book := Like{
		Email_user: likeRequest.Email_user,
		To_Id:      likeRequest.To_id,
		Post_user:  likeRequest.Post_user,
		Is_Like:    likeRequest.Is_Like,
	}

	newBook, err := s.repository.Create(book)

	return newBook, err
	// return s.repository.FindAll()
}

func (s *service) FindByUserLike(id string, email_user string) ([]Like, error) {
	books, err := s.repository.FindByUserLike(id, email_user)
	return books, err
}
	
