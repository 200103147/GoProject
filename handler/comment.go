package handler

import (
	"fmt"
	"net/http"

	"github.com/achmadrizkin/go_social_media_API/comment"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type commentHandler struct {
	commentService comment.Service
}

func NewCommentHandler(hoodieService comment.Service) *commentHandler {
	return &commentHandler{hoodieService}
}

func (h *commentHandler) PostCommentHandler(c *gin.Context) {
	var hoodieRequest comment.CommentRequest

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

	hoodie, err := h.commentService.Create(hoodieRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": hoodie,
	})
}

func (h *commentHandler) GetToUserAndPostUser(c *gin.Context) {
	postUser := c.Param("PostUser")
	toId := c.Param("toId")

	comments, err := h.commentService.GetToUserAndPostUser(postUser, toId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	var commentsResponse []comment.CommentResponse

	for _, b := range comments {
		commentResponse := converToCommentResponse(b)
		commentsResponse = append(commentsResponse, commentResponse)
	}

	//BEWARE DONT TOUCH THIS CODE
	if commentsResponse != nil {
		c.JSON(http.StatusOK, gin.H{
			"data": commentsResponse,
		})
	}
}

func converToCommentResponse(b comment.Comment) comment.CommentResponse {
	return comment.CommentResponse{
		Id:         b.Id,
		ToId:       b.ToId,
		Name_user:  b.Name_user,
		Email_user: b.Email_user,
		Comment:    b.Comment,
		PostUser:   b.PostUser,
		ToUser:     b.ToUser,
		Image_url:  b.Image_url,
	}
}
