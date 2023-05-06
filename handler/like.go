package handler

import (
	"fmt"
	"net/http"

	"github.com/achmadrizkin/go_social_media_API/like"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type likeHandler struct {
	likeService like.Service
}

func NewLikeHandler(hoodieService like.Service) *likeHandler {
	return &likeHandler{hoodieService}
}

func (h *likeHandler) PostLikeHandler(c *gin.Context) {
	var hoodieRequest like.LikeRequest

	err := c.ShouldBindJSON(&hoodieRequest)

	if err != nil {
		// log.Fatal(err) -> kalau terjadi error, server mati
		for _, e := range err.(validator.ValidationErrors) {
			errMessage := fmt.Sprintf("Error on filled %s, condition: %s", e.Field(), e.ActualTag())
			c.JSON(http.StatusBadRequest, errMessage)

			// gunakan return untuk tidak melanjutkan yang dibawah
			return
		}
	}

	hoodie, err := h.likeService.Create(hoodieRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": hoodie,
	})
}

func (h *likeHandler) GetLikeByUser(c *gin.Context) {
	id := c.Param("id")
	email_user := c.Param("email_user")

	allproductss, err := h.likeService.FindByUserLike(id, email_user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	var allproductssResponse []like.LikeResponse

	for _, b := range allproductss {
		allproductsResponse := converToLikeResponse(b)
		allproductssResponse = append(allproductssResponse, allproductsResponse)
	}

	//BEWARE DONT TOUCH THIS CODE
	if allproductssResponse != nil {
		c.JSON(http.StatusOK, gin.H{
			"data": allproductssResponse,
		})
	}
}

func converToLikeResponse(b like.Like) like.LikeResponse {
	return like.LikeResponse {
		Email_user: b.Email_user,
		To_id:      b.To_Id,
		Post_user:  b.Post_user,
		Is_Like:    b.Is_Like,
	}
}
