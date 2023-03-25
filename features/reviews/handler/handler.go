package handler

import (
	"log"
	"musiclab-be/features/reviews"
	"musiclab-be/utils/helper"
	"net/http"

	"github.com/labstack/echo/v4"
)

type reviewControll struct {
	srv reviews.ReviewService
}

func New(srv reviews.ReviewService) reviews.ReviewHandler {
	return &reviewControll{
		srv: srv,
	}
}

// PostMentorReview implements reviews.ReviewHandler
func (rc *reviewControll) PostMentorReview() echo.HandlerFunc {
	return func(c echo.Context) error {
		studentID := helper.ExtractTokenUserId(c)
		input := AddMentorReview{}
		err := c.Bind(&input)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "input format incorrect"})
		}
		mentorGenre := addMentorReviewToCore(input)
		mentorGenre.StudentID = studentID
		err = rc.srv.PostMentorReview(mentorGenre)
		if err != nil {
			log.Println("error running add mentor genre service: ", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "server problem"})
		}
		return c.JSON(http.StatusCreated, map[string]interface{}{
			"message": "success make a review",
		})
	}
}

// GetMentorReview implements reviews.ReviewHandler
func (*reviewControll) GetMentorReview() echo.HandlerFunc {
	panic("unimplemented")
}
