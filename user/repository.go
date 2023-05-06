package user

import (
	"fmt"

	"gorm.io/gorm"
)

type Repository interface {
	FindByNameProduct(name_product string) ([]User, error)
	FindByEmail(name_product string) ([]User, error)
	FindByID(ID int) (User, error)
	Update(user User) (User, error)
	CreateIfNotExistOrUpdateIfExist(email string) ([]User, error)
	UpdatePost(email string, post int) ([]User, error)
	UpdateUserFollowing(email string, post int) ([]User, error)
	UpdateUserFollowers(email string, post int) ([]User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindByNameProduct(name_user string) ([]User, error) {
	var user []User
	// err := r.db.Where("title = ?", title).First(&users).Error

    value := fmt.Sprintf("%%%s%%", name_user)
	err := r.db.Where("name_user LIKE ?", value).Find(&user).Error

	// err := r.db.Where("email_user LIKE ?", name_product).Find(&allProducts).Error

	return user, err
}

func (r *repository) FindByEmail(email string) ([]User, error) {
	var user []User
	// err := r.db.Where("title = ?", title).First(&users).Error

	err := r.db.Where("email_user LIKE ?", email).Find(&user).Error

	// err := r.db.Where("email_user LIKE ?", name_product).Find(&allProducts).Error

	return user, err
}

func (r *repository) CreateIfNotExistOrUpdateIfExist(email string) ([]User, error) {
	var user []User 
	err := r.db.Where(User{Email_user: email}).Assign(User{Email_user: email}).FirstOrCreate(&User{}).Error

	return user, err
}

func (r *repository) UpdatePost(email string, post int) ([]User, error) {
	var user []User 
	err := r.db.Where(User{Email_user: email}).Assign(User{Email_user: email, Post: post}).FirstOrCreate(&User{}).Error

	return user, err
}

func (r *repository) UpdateUserFollowing(email string, post int) ([]User, error) {
	var user []User 
	err := r.db.Where(User{Email_user: email}).Assign(User{Email_user: email, Following: post}).FirstOrCreate(&User{}).Error

	return user, err
}

func (r *repository) UpdateUserFollowers(email string, post int) ([]User, error) {
	var user []User 
	err := r.db.Where(User{Email_user: email}).Assign(User{Email_user: email, Followers: post}).FirstOrCreate(&User{}).Error

	return user, err
}

func (r *repository) FindByID(ID int) (User, error) {
	var allProducts User
	err := r.db.Find(&allProducts, ID).Error

	return allProducts, err
}

func (r *repository) Update(allProduct User) (User, error) {
	err := r.db.Save(&allProduct).Error

	return allProduct, err
}