package comment

type Service interface {
	Create(commentRequest CommentRequest) (Comment, error)
	GetToUserAndPostUser(post_user string, to_id string) ([]Comment, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Create(commentRequest CommentRequest) (Comment, error) {
	book := Comment {
		Name_user: commentRequest.Name_user,
		ToId: commentRequest.ToId,
		Email_user: commentRequest.Email_user,
		Comment: commentRequest.Comment,
		ToUser: commentRequest.ToUser,
		PostUser: commentRequest.PostUser,
		Image_url: commentRequest.Image_url,
	}

	newBook, err := s.repository.Create(book)

	return newBook, err
	// return s.repository.FindAll()
}

func (s *service) GetToUserAndPostUser(post_user string, to_id string) ([]Comment, error) {
	books, err := s.repository.GetToUserAndPostUser(post_user, to_id)
	return books, err
}