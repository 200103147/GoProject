package handler

import (
	"fmt"
	"net/http"

	"github.com/achmadrizkin/go_social_media_API/reels"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type reelsHandler struct {
	reelsService reels.Service
}

func NewReelsHandler(hoodieService reels.Service) *reelsHandler {
	return &reelsHandler{hoodieService}
}

func (h *reelsHandler) PostReelsHandler(c *gin.Context) {
	var hoodieRequest reels.ReelsRequest

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

	hoodie, err := h.reelsService.Create(hoodieRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": hoodie,
	})
}

func (h *reelsHandler) GetReelsNotUserAndOrderByLike(c *gin.Context) {
	user := c.Param("user")

	allproductss, err := h.reelsService.FindByNotUserAndOrderByLike(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	var allproductssResponse []reels.ReelsResponse

	for _, b := range allproductss {
		allproductsResponse := convertToReelsResponse(b)
		allproductssResponse = append(allproductssResponse, allproductsResponse)
	}

	//BEWARE DONT TOUCH THIS CODE
	if allproductssResponse != nil {
		c.JSON(http.StatusOK, gin.H{
			"data": allproductssResponse,
		})
	}
}

func convertToReelsResponse(b reels.Reels) reels.ReelsResponse {
	return reels.ReelsResponse {
		Id:               b.Id,
		Name_user:        b.Name_user,
		Email_user:       b.Email_user,
		Image_url:        b.Image_url,
		Video_post:       b.Video_post,
		Description_post: b.Description_post,
		Like_post:        b.Like_post,
		CreatedAt:        b.CreatedAt,
		UpdateAt:         b.UpdateAt,
	}
}
