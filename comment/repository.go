package comment

import (
	"fmt"

	"gorm.io/gorm"
)

type Repository interface {
	Create(comment Comment) (Comment, error)
	GetToUserAndPostUser(post_user string, to_id string) ([]Comment, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Create(comment Comment) (Comment, error) {
	err := r.db.Create(&comment).Error

	return comment, err
}

func (r *repository) GetToUserAndPostUser(post_user string, to_id string) ([]Comment, error) {
	var comment []Comment

	value := fmt.Sprintf("%%%s%%", post_user)
	err := r.db.Where("post_user LIKE ? AND to_id LIKE ? ORDER BY created_at DESC", value, to_id).Find(&comment).Error

	return comment, err
}