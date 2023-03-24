package mentors

import (
	"mime/multipart"
	"time"

	"github.com/labstack/echo/v4"
)

type Core struct {
	ID           uint
	AvatarFile   multipart.File
	Avatar       string
	Name         string
	Email        string
	Password     string
	Role         string
	Sex          string
	Phone        string
	Address      string
	Instagram    string
	About        string
	AvgRating    float32
	CountReviews int64
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type MentorsHandler interface {
	GetProfile() echo.HandlerFunc
	GetProfileByIdParam() echo.HandlerFunc
	UpdateData() echo.HandlerFunc
}

type MentorService interface {
	SelectProfile(idMentor uint) (Core, error)
	UpdateData(idMentor uint, input Core) error
}

type MentorData interface {
	SelectProfile(idMentor uint) (Core, error)
	UpdateData(idMentor uint, input Core) error
}
