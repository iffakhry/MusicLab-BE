package services

import (
	"errors"
	"log"
	"musiclab-be/features/reviews"
	"strings"
)

type reviewUseCase struct {
	qry reviews.ReviewData
}

func New(rd reviews.ReviewData) reviews.ReviewService {
	return &reviewUseCase{
		qry: rd,
	}
}

func (ruc *reviewUseCase) PostMentorReview(mentorID uint, newReview reviews.Core) error {
	err := ruc.qry.PostMentorReview(mentorID, newReview)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "data not found"
		} else {
			msg = "server problem"
		}
		log.Println("error add query in service: ", err.Error())
		return errors.New(msg)
	}
	return nil
}

func (ruc *reviewUseCase) GetMentorReview(page, limit int, mentorID uint) ([]reviews.Core, error) {
	offset := (page - 1) * limit

	res, err := ruc.qry.GetMentorReview(limit, offset, mentorID)

	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "review not found"
		} else {
			msg = "there is a problem with the server"
		}
		return []reviews.Core{}, errors.New(msg)
	}

	return res, nil
}
