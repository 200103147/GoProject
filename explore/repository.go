package explore

import (
	"gorm.io/gorm"
)

type Repository interface {
	Create(explore Explore) (Explore, error)
	FindByNotUserAndOrderByLike(user string) ([]Explore, error)
	GetByEmailAndOrderByCreateAt(email string) ([]Explore, error)
	GetByUserFollowing(email_user string) ([]Explore, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Create(explore Explore) (Explore, error) {
	err := r.db.Create(&explore).Error

	return explore, err
}

func (r *repository) FindByNotUserAndOrderByLike(user string) ([]Explore, error) {
	var explore []Explore

	err := r.db.Where("email_user NOT LIKE ? ORDER BY like_post DESC", user).Find(&explore).Error

	return explore, err
}

func (r *repository) GetByEmailAndOrderByCreateAt(email string) ([]Explore, error) {
	var explore []Explore

	err := r.db.Where("email_user LIKE ? ORDER BY created_at DESC", email).Find(&explore).Error

	return explore, err
}

func (r *repository) GetByUserFollowing(email_user string) ([]Explore, error) {
	var explore []Explore

	// err := r.db.Where("email_user LIKE ? ORDER BY created_at DESC", email).Find(&explore).Error
	// err := r.db.Table("user_followers").Select("user_followers.email_user, user_followers.image_url").Joins("inner join explores on explores.email_user = user_followers.user_id").Find(&explore).Table("users").Joins("left join users on users.email_user = userfollowers.user_id").Where("user_followers.email_user = ?", email_user).Find(&user).Error
	err := r.db.Raw("select u.name_user, u.email_user, uf.email_user, ex.image_post,ex.description_post,ex.like_post,uf.image_url, ex.name_user from user_followers uf inner join explores ex on ex.email_user = uf.email_user left join users u on u.email_user = uf.user_id where u.email_user = ?", email_user).Scan(&explore).Error
	return explore, err
}
