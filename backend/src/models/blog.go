package models

import "gorm.io/gorm"

type Blog struct {
    gorm.Model
    BlogID          string `json:"blog_id"`
    BlogName        string `json:"blog_name"`
    BlogTitle       string `json:"blog_title"`
    BlogDescription string `json:"blog_description"`
    BlogImages      string `json:"blog_images"` // Base64 encoded images
}