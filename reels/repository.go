package reels

import (
	"gorm.io/gorm"
)

type Repository interface {
	Create(explore Reels) (Reels, error)
	FindByNotUserAndOrderByLike(user string) ([]Reels, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Create(explore Reels) (Reels, error) {
	err := r.db.Create(&explore).Error

	return explore, err
}

func (r *repository) FindByNotUserAndOrderByLike(user string) ([]Reels, error) {
	var reels []Reels

	err := r.db.Where("email_user NOT LIKE ? ORDER BY like_post DESC", user).Find(&reels).Error

	return reels, err
}