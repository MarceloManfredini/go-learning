package models

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title string // Regular string field
	Body  string // Regular string field
}
