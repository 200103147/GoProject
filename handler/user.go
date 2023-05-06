package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/achmadrizkin/go_social_media_API/user"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type userHandler struct {
	userService user.Service
}

func NewBookHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) GetUserById(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	// call service
	b, err := h.userService.FindByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	allproductsResponse := converToAllUserResponse(b)

	c.JSON(http.StatusOK, gin.H{
		"data": allproductsResponse,
	})
}

func (h *userHandler) GetUserByName(c *gin.Context) {
	name_user := c.Param("name_user")

	b, err := h.userService.FindByNameProduct(name_user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	var booksResponse []user.UserResponse
	for _, b := range b {
		bookResponse := converToAllUserResponse(b)
		booksResponse = append(booksResponse, bookResponse)
	}

	if booksResponse != nil {
		c.JSON(http.StatusOK, gin.H{
			"data": booksResponse,
		})
	}
}

func (h *userHandler) GetUserByEmail(c *gin.Context) {
	email_user := c.Param("email_user")

	b, err := h.userService.FindByEmail(email_user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	var booksResponse []user.UserResponse
	for _, b := range b {
		bookResponse := converToAllUserResponse(b)
		booksResponse = append(booksResponse, bookResponse)
	}

	if booksResponse != nil {
		c.JSON(http.StatusOK, gin.H{
			"data": booksResponse,
		})
	}
}

func (h *userHandler) UpdateUser(c *gin.Context) {
	var u user.UserRequest

	err := c.ShouldBindJSON(&u)

	if err != nil {
		// log.Fatal(err) -> kalau terjadi error, server mati
		for _, e := range err.(validator.ValidationErrors) {
			errMessage := fmt.Sprintf("Error on filled %s, condition: %s", e.Field(), e.ActualTag())
			c.JSON(http.StatusBadRequest, errMessage)

			// gunakan return untuk tidak melanjutkan yang dibawah
			return
		}
	}

	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)
	allproducts, err := h.userService.Update(id, u)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": allproducts,
	})
}

func (h *userHandler) CreateIfNotExistOrUpdateIfExist(c *gin.Context) {
	email_user := c.Param("email_user")

	b, err := h.userService.CreateIfNotExistOrUpdateIfExist(email_user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	var booksResponse []user.UserResponse
	for _, b := range b {
		bookResponse := converToAllUserResponse(b)
		booksResponse = append(booksResponse, bookResponse)
	}

	if booksResponse != nil {
		c.JSON(http.StatusOK, gin.H{
			"data": booksResponse,
		})
	}
}

func (h *userHandler) UpdateUserPost(c *gin.Context) {
	email_user := c.Param("email_user")
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	b, err := h.userService.UpdatePost(email_user, id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	var booksResponse []user.UserResponse
	for _, b := range b {
		bookResponse := converToAllUserResponse(b)
		booksResponse = append(booksResponse, bookResponse)
	}

	if booksResponse != nil {
		c.JSON(http.StatusOK, gin.H{
			"data": booksResponse,
		})
	}
}

func (h *userHandler) UpdateUserFollowing(c *gin.Context) {
	email_user := c.Param("email_user")
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	b, err := h.userService.UpdateUserFollowing(email_user, id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	var booksResponse []user.UserResponse
	for _, b := range b {
		bookResponse := converToAllUserResponse(b)
		booksResponse = append(booksResponse, bookResponse)
	}

	if booksResponse != nil {
		c.JSON(http.StatusOK, gin.H{
			"data": booksResponse,
		})
	}
}

func (h *userHandler) UpdateUserFollowers(c *gin.Context) {
	email_user := c.Param("email_user")
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	b, err := h.userService.UpdateUserFollowers(email_user, id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	var booksResponse []user.UserResponse
	for _, b := range b {
		bookResponse := converToAllUserResponse(b)
		booksResponse = append(booksResponse, bookResponse)
	}

	if booksResponse != nil {
		c.JSON(http.StatusOK, gin.H{
			"data": booksResponse,
		})
	}
}


func converToAllUserResponse(b user.User) user.UserResponse {
	return user.UserResponse{
		Id:         b.Id,
		Name_user:  b.Name_user,
		Email_user: b.Email_user,
		Image_url:  b.Image_url,
		Following:  b.Following,
		Followers:  b.Followers,
		Post:       b.Post,
	}
}
