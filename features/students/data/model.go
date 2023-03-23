package data

import (
	_modelChat "musiclab-be/features/chats/data"
	_modelReview "musiclab-be/features/reviews/data"
	_modelTransaction "musiclab-be/features/transactions/data"

	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Avatar       string
	Name         string `gorm:"type:varchar(50) not null"`
	Email        string `gorm:"not null;unique;type:varchar(50)"`
	Password     string
	Role         string `gorm:"type:varchar(25) not null default 'Student'"`
	Sex          string
	Phone        string `gorm:"type:varchar(12)"`
	Address      string `gorm:"type:varchar(100)"`
	Chats        []_modelChat.Chat
	Reviews      []_modelReview.Review
	Transactions []_modelTransaction.Transaction
}
