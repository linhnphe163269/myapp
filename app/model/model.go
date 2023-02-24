package model

import "github.com/jinzhu/gorm"

type Item struct {
    gorm.Model
    Name        string `json:"name"`
    Description string `json:"description"`
    Price       int    `json:"price"`
}
