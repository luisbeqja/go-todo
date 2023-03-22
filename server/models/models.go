package models

import "gorm.io/gorm"

type TodoItem struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      bool   `json:"status"`
}
