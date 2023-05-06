package like

import (
	"gorm.io/gorm"
)

type Repository interface {
	Create(like Like) (Like, error)
	FindByUserLike(id string, email_user string) ([]Like, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Create(like Like) (Like, error) {
	err := r.db.Create(&like).Error

	return like, err
}

func (r *repository) FindByUserLike(id string, email_user string) ([]Like, error) {
	var like []Like

	err := r.db.Where("to_id LIKE ? AND email_user LIKE ?", id, email_user).Find(&like).Error

	return like, err
}
