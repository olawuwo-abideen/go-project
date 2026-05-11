package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Post struct {
	ID        string         `json:"id" gorm:"type:char(36);primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" swaggerignore:"true"`

	Title string `json:"title"`
	Body  string `json:"body"`
}

func (post *Post) BeforeCreate(tx *gorm.DB) (err error) {
	post.ID = uuid.New().String()
	return
}
