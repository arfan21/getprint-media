package models

import (
	"time"

	"gopkg.in/guregu/null.v4"
)

type Media struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt null.Time `gorm:"index" json:"deleted_at,omitempty"`
	Url       string    `json:"url"`
	Path      string    `json:"path"`
}
