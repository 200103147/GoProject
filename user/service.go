package user

type Service interface {
	FindByNameProduct(name_product string) ([]User, error)
	FindByEmail(email string) ([]User, error)
	FindByID(ID int) (User, error)
	Update(id int, user UserRequest) (User, error)
	CreateIfNotExistOrUpdateIfExist(email string) ([]User, error)
	UpdatePost(email string, post int) ([]User, error)
	UpdateUserFollowing(email string, post int) ([]User, error)
	UpdateUserFollowers(email string, post int) ([]User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindByNameProduct(name_product string) ([]User, error) {
	users, err := s.repository.FindByNameProduct(name_product)
	return users, err
	// return s.repository.FindAll()
}

func (s *service) FindByEmail(email string) ([]User, error) {
	users, err := s.repository.FindByEmail(email)
	return users, err
	// return s.repository.FindAll()
}

func (s *service) CreateIfNotExistOrUpdateIfExist(email string) ([]User, error) {
	users, err := s.repository.CreateIfNotExistOrUpdateIfExist(email)
	return users, err
	// return s.repository.FindAll()
}

func (s *service) UpdateUserFollowing(email string, post int) ([]User, error) {
	users, err := s.repository.UpdateUserFollowing(email, post)
	return users, err
	// return s.repository.FindAll()
}

func (s *service) UpdateUserFollowers(email string, post int) ([]User, error) {
	users, err := s.repository.UpdateUserFollowers(email, post)
	return users, err
	// return s.repository.FindAll()
}

func (s *service) UpdatePost(email string, post int) ([]User, error) {
	users, err := s.repository.UpdatePost(email, post)
	return users, err
	// return s.repository.FindAll()
}

func (s *service) FindByID(ID int) (User, error) {
	books, err := s.repository.FindByID(ID)
	return books, err
	// return s.repository.FindAll()
}

func (s *service) Update(ID int, u UserRequest) (User, error) {
	user, _ := s.repository.FindByID(ID)

	//
	user.Name_user = u.Name_user
	user.Email_user = u.Email_user
	user.Image_url = u.Image_url

	newAllProduct, err := s.repository.Update(user)

	return newAllProduct, err
	// return s.repository.FindAll()
}